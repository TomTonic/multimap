package art

import "testing"

func TestPresenceBitmap_SetGetClear(t *testing.T) {
	var bm PresenceBitmap

	// check initial state: all bits clear
	for _, i := range []byte{0, 63, 64, 127, 128, 191, 192, 255} {
		if bm.Get(i) {
			t.Fatalf("bit %d should be clear initially", i)
		}
	}

	// set and verify
	for _, i := range []byte{0, 1, 63, 64, 100, 200, 255} {
		bm.Set(i)
		if !bm.Get(i) {
			t.Fatalf("bit %d should be set after Set", i)
		}
	}

	// other bits remain clear
	for _, i := range []byte{2, 62, 65, 199, 254} {
		if bm.Get(i) {
			t.Fatalf("bit %d unexpectedly set", i)
		}
	}

	// clear and verify
	for _, i := range []byte{0, 63, 200, 255} {
		bm.Clear(i)
		if bm.Get(i) {
			t.Fatalf("bit %d should be clear after Clear", i)
		}
	}
}

func TestPresenceBitmap_BulkOperations(t *testing.T) {
	var bm PresenceBitmap

	// Set a contiguous range 50..59
	for i := byte(50); i <= 59; i++ {
		bm.Set(i)
	}
	// verify
	for i := byte(45); i <= 64; i++ {
		want := i >= 50 && i <= 59
		if bm.Get(i) != want {
			t.Fatalf("range check: bit %d expected %v got %v", i, want, bm.Get(i))
		}
	}

	// Clear the whole range
	for i := byte(50); i <= 59; i++ {
		bm.Clear(i)
	}
	for i := byte(50); i <= 59; i++ {
		if bm.Get(i) {
			t.Fatalf("bit %d should be clear after bulk Clear", i)
		}
	}
}
