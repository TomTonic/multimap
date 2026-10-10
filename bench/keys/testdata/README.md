# Real-world key corpora

Package `keys` embeds these files for the key kinds `street`, `path`, `url`
and `links` (`dirs` is derived from `path`).
[`cmd/mkcorpora`](../../cmd/mkcorpora/main.go) downloads their sources from
fixed addresses and rebuilds them byte for byte; the benchmark itself never
touches the network.

## streets.tsv.gz

German street names and the localities that have a street of that name:
425,000 names, 10,199 localities, 1.17 million pairs. The first line is the
number of localities, then one locality per line in ascending order, then one
line per street name in ascending order: the name, a tab, and the
comma-separated indexes of its localities.

Source: [OpenPLZ API data](https://github.com/openpotato/openplzapi.data),
file `src/de/osm/streets.updated.csv` at commit `c26389b`, which is an extract
of the [OpenStreetMap](https://www.openstreetmap.org/) data for Germany.

**License:** © OpenStreetMap contributors. This file is a derived database of
OpenStreetMap data and is made available under the
[Open Database License (ODbL) 1.0](https://opendatacommons.org/licenses/odbl/1-0/),
like its sources. The ODbL applies to this file only, not to the code of this
repository. See <https://www.openstreetmap.org/copyright>.

## paths.txt.gz

600,000 file paths, one per line in ascending order, drawn at random with a
fixed seed from the 7.3 million distinct paths in the packages of Debian 12
"bookworm" (`main`, architectures `all` and `amd64`), as listed in the
archive's `Contents` files on
[snapshot.debian.org](https://snapshot.debian.org/) as of 2026-09-01.

The `Contents` files are part of the freely redistributable Debian archive.
They carry no license of their own and hold only the names of files; this
sample keeps nothing but those names.

## hosts.txt.gz

The 200,000 most popular host names of
[Tranco list 8P48V](https://tranco-list.eu/list/8P48V/1000000), with
subdomains, one per line in rank order. The list was generated on
2026-09-27 from the rankings of Chrome UX Report, Cloudflare Radar,
Farsight, Majestic and Cisco Umbrella of 2026-08-29 to 2026-09-27.
`cmd/mkcorpora` drops the few names with characters other than lowercase
letters, digits, dashes and dots, and names over 64 bytes. The `url` kind
combines these hosts with synthetic paths (see `../url.go`).

Tranco: Victor Le Pochat, Tom Van Goethem, Samaneh Tajalizadehkhoob, Maciej
Korczyński and Wouter Joosen, "Tranco: A Research-Oriented Top Sites Ranking
Hardened Against Manipulation", *Proceedings of the 26th Network and
Distributed System Security Symposium (NDSS 2019)*.

**License:** the list combines its sources' data under their terms: Cisco
Umbrella free of charge, Majestic under CC BY 3.0, the Chrome UX Report
under CC BY-SA 4.0 and Cloudflare Radar under
[CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/). This file is
therefore made available under CC BY-NC 4.0, for non-commercial use such as
this benchmark. The license applies to this file only, not to the code of
this repository.

## links.tsv.gz

The links between the articles of Simple English Wikipedia: 140,000 pages and
the 6.5 million links from them to 974,425 distinct target titles. The first
line is the number of target titles, then one target title per line in
ascending order, then one line per page in ascending order: the page's title, a
tab and the comma-separated indexes of the titles it links to, ascending, each
counted from 0 in the list of target titles. Titles are as stored in the
database: UTF-8, with underscores for spaces. The file is plain text with
absolute indexes, in the format of `streets.tsv.gz`, so that both are read by
the same code.

Source: the database dump of Simple English Wikipedia of 2026-10-01,
<https://dumps.wikimedia.org/simplewiki/20261001/>, the files
`simplewiki-20261001-page.sql.gz` (SHA-1 `765d7db7…`),
`simplewiki-20261001-linktarget.sql.gz` (`1f3017c5…`) and
`simplewiki-20261001-pagelinks.sql.gz` (`b6795baa…`). The links of a page are
the rows of `pagelinks` with `pl_from` = the page and a `linktarget` of
namespace 0; only pages of namespace 0 (articles and redirects) are keys, and
only the pages with at least one such link: 399,039 of them. `cmd/mkcorpora`
draws 140,000 of those at random with a fixed seed. A page that is drawn keeps
all its links, so the number of links a page has is distributed as in the whole
wiki; the targets are those of the drawn pages and make a part of the wiki's
1.7 million target titles. Targets include titles that have no page (red
links): 719,666 of the 974,425, as in the dump, where 27% of the links lead to
a title that does not exist.

|  | this file | the whole wiki (399,039 pages) |
|---|--:|--:|
| pages, links | 140,000, 6,496,745 | 399,039, 18,277,816 |
| links a page: mean, median, max | 46.4, 11, 5,693 | 45.8, 11, 5,693 |
| pages with 1 / 2-8 / 9-64 / 65+ links | 29.0 / 16.1 / 38.2 / 16.7 % | 29.2 / 16.1 / 38.3 / 16.5 % |
| links in pages with 65+ links | 77.2 % | 76.9 % |
| redirects among the pages | 28.7 % | 28.8 % |
| length of a page title: mean | 16.7 B (1 to 255) | 16.7 B |

The pages with one link are almost all redirects: in the whole wiki 113,943 of
the 116,328 pages with one link (98 %) are redirects, and 99.2 % of the redirects
have exactly one link.
By number of links: 1: 29.0 %, 2: 1.4 %, 3-4: 4.6 %, 5-8: 10.1 %, 9-16: 14.3 %, 17-64: 23.9 %, 65-256: 12.5 %, 257 and
more: 4.2 % of the pages; the links by the same classes: 0.6, 0.1, 0.3, 1.4, 3.7, 16.6, 33.8, 43.5 %.
`cmd/mkcorpora links` prints these numbers (it downloads 155 MB, needs about
a minute and 1 GB of memory, and gives the same file every time with the same Go version).

**License:** the text of Wikipedia is available under the
[Creative Commons Attribution-ShareAlike 4.0 License](https://creativecommons.org/licenses/by-sa/4.0/)
(and the GNU Free Documentation License), see
<https://dumps.wikimedia.org/legal.html> and section 7 of the
[Wikimedia Foundation Terms of Use](https://foundation.wikimedia.org/wiki/Policy:Terms_of_Use).
The titles and the links between them are names and facts, for which the
contributors have waived their database rights, but this file is made available
under CC BY-SA 4.0 as a precaution, like its source. The license applies to this
file only, not to the code of this repository. Attribution: Wikipedia
contributors, [Simple English Wikipedia](https://simple.wikipedia.org/), whose
page histories list the authors of every article; this file is a selection of
titles and links from the dump named above, without changes to the titles.
