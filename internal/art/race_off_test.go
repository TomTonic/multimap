//go:build !race

package art

// underRace is true when the tests run with the race detector (see race_on_test.go).
const underRace = false
