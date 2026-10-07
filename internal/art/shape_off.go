//go:build !mkstats

package art

// Shape is empty without the build tag mkstats (see shape_on.go).
func (m *Map[T]) Shape() string { return "" }
