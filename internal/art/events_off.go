//go:build !mkstats

package art

// event names what happens to the multi-key pages of a tree; see events_on.go. Without the
// build tag mkstats nothing is counted, and the calls cost nothing.
type event uint8

const (
	evPair event = iota
	evPairNo
	evAdded
	evPromote
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

// ev counts event e once, with n entries of the page it concerns.
func ev(event, int) {}

// EventsEnabled says whether this build counts the events of the multi-key pages.
const EventsEnabled = false

// Events returns the counts of the events and their histograms over the entries of the page;
// nil in a build without the tag mkstats.
func Events() map[string]*EventCount { return nil }

// ResetEvents sets all counts to zero.
func ResetEvents() {}

// EventCount is how often an event happened, and how often it concerned a page of n entries.
type EventCount struct {
	Total   uint64
	Entries [257]uint64
}
