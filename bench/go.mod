module github.com/TomTonic/multimap/bench

go 1.27.1

replace github.com/TomTonic/multimap => ../

require (
	github.com/TomTonic/Set3 v0.4.2
	github.com/TomTonic/multimap v0.0.0-00010101000000-000000000000
	github.com/TomTonic/rtcompare v0.6.0
	github.com/plar/go-adaptive-radix-tree/v2 v2.0.4
	github.com/plar/go-hot-trie v0.1.2
	github.com/tidwall/btree v1.8.1
)

require (
	github.com/dolthub/maphash v0.1.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
	golang.org/x/text v0.42.0 // indirect
)
