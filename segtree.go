package segtree

type SegTree[T any] struct {
	function func(values ...T) T
	array    []*T
	root     node[T]
}

func (s *SegTree[T]) Setup(arr []T, f func(values ...T) T) {
	s.function = f
	s.array = make([]*T, len(arr))

	s.root.setup(arr, s.array, s.function)
}

func (s *SegTree[T]) Update(val T, index int) {
	s.root.update(val, index, s.function)
}

func (s *SegTree[T]) UpdateRange(arr []T, start int, end int) {
	s.root.insert(arr, start, end, s.function)
}

func (s *SegTree[T]) Query(left int, right int) T {
	result := s.root.query(left, right, s.function)

	return result
}

func (s *SegTree[T]) GetArray() []T {
	values := make([]T, len(s.array))
	for i, p := range s.array {
		values[i] = *p
	}

	return values
}
