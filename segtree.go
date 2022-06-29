package segtree

// SegTree structure contains field required to work with generic
// segment trees
type SegTree[T any] struct {
	function func(values ...T) T
	array    []*T
	root     node[T]
}

// The method Setup() lets build a generic segment trees
// It requires a type to be specified, a slice of elements to build the tree on
// and a function to be calculated for each segment
func (s *SegTree[T]) Setup(arr []T, f func(values ...T) T) {
	s.function = f
	s.array = make([]*T, len(arr))

	s.root.setup(arr, s.array, s.function)
}

// The method Update() allows to update a singular element of the array
// with subsequent updating of nodes
func (s *SegTree[T]) Update(val T, index int) {
	s.root.update(val, index, s.function)
}

// The method UpdateRange() allow to update a subslice of the array
// with subsequent updating of nodes
func (s *SegTree[T]) UpdateRange(arr []T, start int, end int) {
	s.root.insert(arr, start, end, s.function)
}

// The method Query() allows to calculate a query over a specified segment
// For instance, if a segment tree is built from a slice []int{1, 2, 3, 4, 5}
// to calculate sum of elements on a segment, executing Query(1, 3) will return 9 (2+3+4)
func (s *SegTree[T]) Query(left int, right int) T {
	result := s.root.query(left, right, s.function)

	return result
}

// The method GetArray() allows to get a slice from which a segment tree is build
func (s *SegTree[T]) GetArray() []T {
	values := make([]T, len(s.array))
	for i, p := range s.array {
		values[i] = *p
	}

	return values
}
