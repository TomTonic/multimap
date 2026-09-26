# Real-world key corpora

Package `keys` embeds these files for the key kinds `street` and `path`.
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
