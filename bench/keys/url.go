package keys

import (
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/TomTonic/rtcompare"
)

// The url kind draws real host names from the Tranco list
// (testdata/hosts.txt.gz, see cmd/mkcorpora) and gives each host paths shaped
// like those of a real site. The hosts, and not the paths, decide how the
// keys branch near their start, which is where a synthetic corpus of a few
// hosts differs most from the web.
//
// A key's host is drawn by rank, with weight 1/(rank + hostSkew): popular
// hosts hold more addresses, as in a crawl or a link index, but the top host
// holds only about 1% of them. Every host has a fixed profile, drawn from its
// rank alone: whether it is served under "www.", whether its paths start
// with a language, which of the path styles below it mostly uses, and the
// few sections its paths start with. A key uses the host's own style 70% of
// the time and another style otherwise; a quarter of the keys carry a query.

// hostSkew flattens the head of the host distribution (see above).
const hostSkew = 10

// Path styles of a site.
const (
	styleArticle = iota // /news/2024/03/some-long-slug.html
	styleShop           // /shoes/boots/some-slug-p123456
	styleStatic         // /img/3f/a9/3fa9…c1.jpg
	styleAPI            // /api/v2/orders/123456/items
	styleWiki           // /en/wiki/Some_Title_Words
	numStyles
)

// site is the fixed profile of one host.
type site struct {
	host     string // with "www." if the site is served there
	langs    []string
	style    int
	sections []string
}

// urlModel holds what every url generator shares: the hosts with their
// weights and profiles, and a vocabulary of path words with its weights.
type urlModel struct {
	hosts []string
	hostW zipf
	sites []*site // profiles, made on first use
	mu    sync.Mutex
	vocab []string
	wordW zipf
}

var theURLModel = sync.OnceValue(func() *urlModel {
	hosts := hostCorpus()
	rng := rtcompare.NewDPRNG(0x75726c) // fixed: the vocabulary is part of the kind
	vocab := make([]string, 0, 4096)
	seen := map[string]bool{}
	for len(vocab) < cap(vocab) {
		if w := pathWord(&rng); !seen[w] {
			seen[w] = true
			vocab = append(vocab, w)
		}
	}
	return &urlModel{
		hosts: hosts, hostW: newZipf(len(hosts), hostSkew), sites: make([]*site, len(hosts)),
		vocab: vocab, wordW: newZipf(len(vocab), 1),
	}
})

// pathWord returns a pronounceable lowercase word of one to four syllables.
func pathWord(rng *rtcompare.DPRNG) string {
	const cons, vows = "bcdfghklmnprstvwz", "aeiou"
	var b strings.Builder
	for range 1 + rng.Uint64()%4 {
		b.WriteByte(cons[rng.Uint64()%uint64(len(cons))])
		b.WriteByte(vows[rng.Uint64()%uint64(len(vows))])
		if rng.Uint64()%3 == 0 {
			b.WriteByte(cons[rng.Uint64()%uint64(len(cons))])
		}
	}
	return b.String()
}

// site returns the profile of the host of rank index i.
func (m *urlModel) site(i int) *site {
	m.mu.Lock()
	defer m.mu.Unlock()
	if s := m.sites[i]; s != nil {
		return s
	}
	rng := rtcompare.NewDPRNG(uint64(i) + 1)
	s := &site{host: m.hosts[i], style: int(rng.Uint64() % numStyles)}
	if strings.Count(s.host, ".") == 1 && rng.Uint64()%2 == 0 {
		s.host = "www." + s.host
	}
	if rng.Uint64()%5 == 0 {
		langs := []string{"en", "de", "fr", "es", "it", "ja", "en-us", "en-gb", "de-de", "pt-br"}
		s.langs = langs[:1+rng.Uint64()%uint64(len(langs))]
	}
	for range 2 + rng.Uint64()%9 {
		s.sections = append(s.sections, m.word(&rng))
	}
	m.sites[i] = s
	return s
}

func (m *urlModel) word(rng *rtcompare.DPRNG) string { return m.vocab[m.wordW.draw(rng)] }

// url returns one address.
func (m *urlModel) url(rng *rtcompare.DPRNG) []byte {
	s := m.site(m.hostW.draw(rng))
	k := make([]byte, 0, 128)
	if rng.Uint64()%32 == 0 {
		k = append(k, "http://"...)
	} else {
		k = append(k, "https://"...)
	}
	k = append(append(k, s.host...), '/')
	style := s.style
	if rng.Uint64()%10 >= 7 {
		style = int(rng.Uint64() % numStyles)
	}
	if s.langs != nil && style != styleStatic && style != styleAPI {
		k = append(append(k, s.langs[rng.Uint64()%uint64(len(s.langs))]...), '/')
	}
	section := s.sections[rng.Uint64()%uint64(len(s.sections))]
	switch style {
	case styleArticle:
		k = append(append(k, section...), '/')
		if rng.Uint64()%2 == 0 {
			k = strconv.AppendUint(k, 2005+rng.Uint64()%22, 10)
			mon := 1 + rng.Uint64()%12
			k = append(k, '/', byte('0'+mon/10), byte('0'+mon%10), '/')
		}
		k = m.slug(k, rng, '-')
		if rng.Uint64()%3 == 0 {
			k = append(k, ".html"...)
		}
	case styleShop:
		if rng.Uint64()%4 == 0 {
			k = strconv.AppendUint(append(k, "p/"...), 100000+rng.Uint64()%900000, 10)
			break
		}
		k = append(append(k, section...), '/')
		k = append(append(k, m.word(rng)...), '/')
		k = strconv.AppendUint(append(m.slug(k, rng, '-'), "-p"...), 100000+rng.Uint64()%900000, 10)
	case styleStatic:
		h := rng.Uint64()
		k = append(append(k, section...), '/')
		k = appendHex(k, h>>56, 2)
		k = appendHex(append(k, '/'), h>>48&0xff, 2)
		k = appendHex(append(k, '/'), h, 16)
		exts := []string{".jpg", ".png", ".webp", ".js", ".css", ".svg", ".woff2", ".pdf"}
		k = append(k, exts[rng.Uint64()%uint64(len(exts))]...)
	case styleAPI:
		k = strconv.AppendUint(append(k, "api/v"...), 1+rng.Uint64()%3, 10)
		k = append(append(append(k, '/'), section...), '/')
		k = strconv.AppendUint(k, rng.Uint64()%10_000_000, 10)
		if rng.Uint64()%2 == 0 {
			k = append(append(k, '/'), m.word(rng)...)
		}
	case styleWiki:
		k = append(k, "wiki/"...)
		for i := range 1 + rng.Uint64()%4 {
			if i > 0 {
				k = append(k, '_')
			}
			w := m.word(rng)
			k = append(append(k, w[0]-'a'+'A'), w[1:]...)
		}
	}
	if rng.Uint64()%4 == 0 {
		k = m.query(k, rng)
	}
	return k
}

// slug appends two to seven words joined by sep, as in a title.
func (m *urlModel) slug(k []byte, rng *rtcompare.DPRNG, sep byte) []byte {
	for i := range 2 + rng.Uint64()%6 {
		if i > 0 {
			k = append(k, sep)
		}
		k = append(k, m.word(rng)...)
	}
	return k
}

// query appends one to three distinct parameters. They hold no slash, so
// that a key's directory (see Prefix) ends before its query.
func (m *urlModel) query(k []byte, rng *rtcompare.DPRNG) []byte {
	var used uint8
	for i := range 1 + rng.Uint64()%3 {
		p := rng.Uint64() % 6
		for used&(1<<p) != 0 {
			p = (p + 1) % 6
		}
		used |= 1 << p
		k = append(k, "?&"[min(i, 1)])
		switch p {
		case 0:
			k = strconv.AppendUint(append(k, "id="...), rng.Uint64()%1_000_000, 10)
		case 1:
			k = strconv.AppendUint(append(k, "page="...), 1+rng.Uint64()%50, 10)
		case 2:
			k = append(append(append(append(k, "q="...), m.word(rng)...), '+'), m.word(rng)...)
		case 3:
			k = append(append(k, "utm_source="...), m.word(rng)...)
		case 4:
			k = appendHex(append(k, "ref="...), rng.Uint64(), 8)
		default:
			k = strconv.AppendUint(append(k, "v="...), rng.Uint64()%100, 10)
		}
	}
	return k
}

// appendHex appends the low n hex digits of v.
func appendHex(k []byte, v uint64, n int) []byte {
	const hex = "0123456789abcdef"
	for i := n - 1; i >= 0; i-- {
		k = append(k, hex[v>>(4*i)&0xf])
	}
	return k
}

// zipf draws indexes 0..n-1 with weight 1/(i + 1 + q).
type zipf struct{ cum []float64 }

func newZipf(n int, q float64) zipf {
	cum := make([]float64, n)
	s := 0.0
	for i := range cum {
		s += 1 / (float64(i+1) + q)
		cum[i] = s
	}
	return zipf{cum}
}

func (z zipf) draw(rng *rtcompare.DPRNG) int {
	u := rng.Float64() * z.cum[len(z.cum)-1]
	i, _ := slices.BinarySearch(z.cum, u)
	return min(i, len(z.cum)-1)
}
