# Generic segment tree implementation

This package implements a segment tree with nodes of a generic type. Segment tree is a data structure allowing to
calculate a function on a segment of a slice in *O(log n)* time. More about segment tree: https://en.wikipedia.org/wiki/Segment_tree.

## Installation
Since the package uses generics, the version of Go must be >= 1.18

Installing:
```
go get github.com/kwarabei/segtree
```

## API

The package has a public struct called SegTree with a defined set of methods.

Building a segment tree (the built tree will be used across next examples). You should pass an arbitrary function to be calculated while building the tree.
```go
var s SegTree[int]
 sumFunc := func(values ...int) int {
	s := 0
	for _, v := range values {
		s += v
	}

	return s
}

s.Setup([]int{1, 2, 3, 4, 5}, sumFunc)
```

Calculating a query (should equal 9). Starting and ending indices should be specified inclusively.
```go
result := s.Query(0, 2)
```

Updating a single element. The first argument of the method is a new value, the second is an index. 
```go
s.Update(10, 3)
```

Updating a range of element. Indices should be specified inclusively.
```go
s.UpdateRange([]int{1, 1}, 2, 3)
```

Retrieving a slice of elements from which the tree is built.
```go
values := s.GetArray()
fmt.Printf("%+v\n", values)
```

The test file of the project contains a few more examples of segment tree usage for different types.
