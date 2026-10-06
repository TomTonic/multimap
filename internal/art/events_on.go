//go:build mkstats

package art

// event names what happens to the multi-key pages of a tree. A build with the tag mkstats
// counts them (all trees of the process together, not safe for concurrent use), so that
// bench/cmd/bench's probe can tell what a workload does to the pages (docs/redesign/step4-probe.md).
type event uint8

const (
	evPair event = iota
	evPairNo
	evAdded
	evAddedValue
	evBurst
	evWiden
	evAbove
	evPageRemove
	evToSingleKey
	evShrink
	evMergeTry
	evMergeOK
	evPrependNo
	evBuildPage
	evBuildNode
	evBuildSingleKey
	evCount
)

var eventNames = [evCount]string{
	evPair:           "pair: single-key page and new key make a page",
	evPairNo:         "pair refused: the two do not make a page",
	evAdded:          "key added to a page in place",
	evAddedValue:     "value added to a key in a page in place",
	evBurst:          "burst: page is full, page dissolves",
	evWiden:          "widen: key outside the common prefix joins the page",
	evAbove:          "byte node put above a page",
	evPageRemove:     "entry removed from a page",
	evToSingleKey:    "page with one entry left becomes a single-key page",
	evShrink:         "page shrinks to a smaller class",
	evMergeTry:       "merge tried (node with children that are pages)",
	evMergeOK:        "merge done (node becomes a page)",
	evPrependNo:      "prepend refused: page stays below the node",
	evBuildPage:      "build: page made",
	evBuildNode:      "build: byte node made",
	evBuildSingleKey: "build: single-key page made",
}

// EventsEnabled says whether this build counts the events of the multi-key pages.
const EventsEnabled = true

// EventCount is how often an event happened, and how often it concerned a page of n entries
// (n above 255 counts as 256).
type EventCount struct {
	Total   uint64
	Entries [257]uint64
}

var counts [evCount]EventCount

// ev counts event e once, with n entries of the page it concerns.
func ev(e event, n int) {
	counts[e].Total++
	counts[e].Entries[min(max(n, 0), 256)]++
}

// Events returns the counts of the events and their histograms over the entries of the page.
func Events() map[string]*EventCount {
	out := make(map[string]*EventCount, evCount)
	for e := range counts {
		c := counts[e]
		out[eventNames[e]] = &c
	}
	return out
}

// ResetEvents sets all counts to zero.
func ResetEvents() { counts = [evCount]EventCount{} }
