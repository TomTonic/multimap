package art

import set3 "github.com/TomTonic/Set3"

// NOTE:
// - 64-bit architecture assumed (8-byte pointer width).
// - All node types below use explicit padding bytes to keep child arrays
//   aligned to 8‑byte boundaries and/or to hit exact 2^n total sizes.
// - Generic parameter [T comparable] is used for typed payload pointers.
//
// Reasons for the layout and padding choices:
// - Cache-line optimization: node sizes are chosen so nodes align with
//   common cache-line boundaries (and often fit in a small number of
//   consecutive cache lines). This reduces cache misses and improves
//   prefetch behaviour for tree traversals.
// - Go allocation behaviour: allocating objects that map well to the
//   runtime allocator's size-classes reduces heap fragmentation and
//   the amount of wasted space; keeping node objects compact helps
//   the garbage collector and allocator operate more efficiently.
// - Memory fragmentation: fixed-size, power-of-two-friendly objects
//   are less likely to fragment memory under long-running workloads.
//
// Padding explanation and GC considerations:
// - Explicit padding ensures child arrays and pointer fields are
//   8-byte aligned which avoids unaligned memory accesses and
//   improves pointer load/store performance on 64-bit CPUs.
// - The `value`/child pointer placement is chosen so that pointer
//   words fall on predictable offsets — this helps the GC scanner
//   (reduces false positives and keeps the live-pointer layout stable).
// - Padding bytes are non-pointer data and do not affect GC scans; by
//   keeping pointer fields at aligned offsets we avoid extra per-object
//   scan overhead.

const (
	maxLocalPrefixLen = 14

	maxChildrenLeaf     = 1
	maxChildrenNode64   = 4
	maxChildrenNode128  = 11
	maxChildrenNode256  = 25
	maxChildrenNode512  = 50
	maxChildrenNode1024 = 107
	maxChildrenFullNode = 256
)

// -----------------------------------------------------------------------------
// Common base (header + value), 24 bytes total.
//
// Layout:
//
//	meta        : 1 byte  (high nibble = node kind, low nibble = inline prefix length)
//	numChildren : 1 byte  (number of children, 0..k)
//	localPrefix : 14 bytes (inline prefix payload)
//	value       : 8 bytes  (*set3.Set3[T])
//
// -----------------------------------------------------------------------------
type Node[T comparable] struct {
	meta        uint8
	numChildren uint8
	localPrefix [maxLocalPrefixLen]byte
	value       *set3.Set3[T] // 8 B at offset 16 -> GC-scan friendly
	// Total: 24 B
}

// -----------------------------------------------------------------------------
// LeafNode — EXACT 32 bytes
// Specialized leaf node holding a single child pointer.
//
// Layout:
//
//	Node (24) + child pointer (8) = 32
//
// -----------------------------------------------------------------------------
type LeafNode[T comparable] struct {
	Node[T]
	child *Node[T]
	// Total: 32 B
}

// -----------------------------------------------------------------------------
// Node64 — EXACT 64 bytes
// No bitmap; UNSORTED; linear/unrolled lookup.
// Size calculation:
//
//	Node (24) + firstKeyByte[4] (4) + pad4 (4) + child[4]*8 (32) = 64
//
// -----------------------------------------------------------------------------
type Node64[T comparable] struct {
	Node[T]
	firstKeyByte [maxChildrenNode64]byte
	pad4         [4]byte                     // pad to 8-byte boundary before child[]
	child        [maxChildrenNode64]*Node[T] // 32 B at offset 32 -> GC-scan friendly
	// Total: 64 B
}

// -----------------------------------------------------------------------------
// Node128 — EXACT 128 bytes
// No bitmap; UNSORTED; linear/unrolled lookup.
// Size calculation:
//
//	Node (24) + firstKeyByte[11] (11) + pad5 (5) + child[11]*8 (88) = 128
//
// -----------------------------------------------------------------------------
type Node128[T comparable] struct {
	Node[T]
	firstKeyByte [maxChildrenNode128]byte
	pad5         [5]byte                      // pad to 8-byte boundary before child[]
	child        [maxChildrenNode128]*Node[T] // 88 B at offset 40 -> GC-scan friendly
	// Total: 128 B
}

// -----------------------------------------------------------------------------
// Node256 — EXACT 256 bytes
// No bitmap; UNSORTED; linear/unrolled lookup.
// Size calculation:
//
//	Node (24) + firstKeyByte[25] (25) -> 49
//	+ pad7 (7) -> align to 8-byte boundary (56)
//	+ child[25]*8 (200) = 256
//
// -----------------------------------------------------------------------------
type Node256[T comparable] struct {
	Node[T]
	firstKeyByte [maxChildrenNode256]byte
	pad7         [7]byte                      // pad to 8-byte boundary before child[]
	child        [maxChildrenNode256]*Node[T] // 200 B at offset 56 -> GC-scan friendly
	// Total: 256 B
}

// -----------------------------------------------------------------------------
// Node512 — EXACT 512 bytes
// Bitmap present; SORTED firstKeyByte[]; binary search anticipated.
// Size calculation:
//
//	Node (24) + bitmap (32) = 56
//	+ firstKeyByte[50] (50) -> 106
//	+ pad6 (6) -> align to 8-byte boundary (112)
//	+ child[50]*8 (400) = 512
//
// -----------------------------------------------------------------------------
type Node512[T comparable] struct {
	Node[T]
	bitmap       PresenceBitmap
	firstKeyByte [maxChildrenNode512]byte     // SORTED
	pad6         [6]byte                      // pad to 8-byte boundary before child[]
	child        [maxChildrenNode512]*Node[T] // 400 B at offset 112 -> GC-scan friendly
	// Total: 512 B
}

// -----------------------------------------------------------------------------
// Node1024 — EXACT 1024 bytes
// Bitmap present; SORTED firstKeyByte[]; binary search anticipated.
// Size calculation:
//
//	Node (24) + bitmap (32) = 56
//	+ firstKeyByte[107] (107) -> 163
//	+ pad5 (5) -> align to 8-byte boundary (168)
//	+ child[107]*8 (856) = 1024
//
// -----------------------------------------------------------------------------
type Node1024[T comparable] struct {
	Node[T]
	bitmap       PresenceBitmap
	firstKeyByte [maxChildrenNode1024]byte     // SORTED
	pad5         [5]byte                       // pad to 8-byte boundary before child[]
	child        [maxChildrenNode1024]*Node[T] // 856 B at offset 168 -> GC-scan friendly
	// Total: 1024 B
}

// -----------------------------------------------------------------------------
// FullNode — EXACT 64 B node object + EXACT 2048 B external array
// Bitmap present; direct index into external child[256]; no sorting required.
// Node object layout:
//
//	Node (24) + bitmap (32) + array pointer (8) = 64
//
// External array layout:
//
//	[256]*Node[T] = 256 * 8 = 2048 bytes
//
// -----------------------------------------------------------------------------
type FullNode[T comparable] struct {
	Node[T]
	bitmap PresenceBitmap
	child  *[maxChildrenFullNode]*Node[T] // external array (2048 B)
	// Total (node object): 64 B (exact)
}
