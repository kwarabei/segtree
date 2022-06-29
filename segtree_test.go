package segtree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntSumSegTree(t *testing.T) {
	var s SegTree[int]
	sumFunc := func(values ...int) int {
		s := 0
		for _, v := range values {
			s += v
		}

		return s
	}

	s.Setup([]int{5, 3, 4, 2, 1}, sumFunc)
	queries := [][]int{{0, 4}, {1, 3}, {2, 4}}
	expectedResults := []int{15, 9, 7}
	for i := range queries {
		start := queries[i][0]
		end := queries[i][1]
		res := s.Query(start, end)
		if res != expectedResults[i] {
			errMsg := fmt.Sprintf("Calculated sum query for elements from %d to %d"+
				"(inclusively) of an array %+v\n"+
				"Expected result: %d\n"+
				"Got instead: %d\n", start, end, s.GetArray(), expectedResults[i], res)

			t.Fatalf(errMsg)
		}

	}
}

func TestFlaotMinSegTree(t *testing.T) {
	var s SegTree[float64]
	minFunc := func(values ...float64) float64 {
		var min float64 = values[0]
		for i := 1; i < len(values); i++ {
			if values[i] < min {
				min = values[i]
			}
		}

		return min
	}

	s.Setup([]float64{4.0, 1.5, 15.032, 8.66, 0.101, 15.16, 3.005}, minFunc)
	queries := [][]int{{0, 6}, {2, 3}, {3, 5}}
	expectedResults := []float64{0.101, 8.66, 0.101}
	for i := range queries {
		start := queries[i][0]
		end := queries[i][1]
		res := s.Query(start, end)
		if res != expectedResults[i] {
			errMsg := fmt.Sprintf("Calculated min query for elements from %d to %d"+
				"(inclusively) of an array %+v\n"+
				"Expected result: %f\n"+
				"Got instead: %f\n", start, end, s.GetArray(), expectedResults[i], res)

			t.Fatalf(errMsg)
		}
	}
}

func TestUpdatesSegTree(t *testing.T) {
	var s SegTree[int]
	mulFunc := func(values ...int) int {
		mul := 1
		for _, v := range values {
			mul *= v
		}

		return mul
	}

	s.Setup([]int{2, 3, 7, 2, 6, 4}, mulFunc)
	q := s.Query(0, 2)
	if q != 42 {
		t.Fatalf("Failed to calculate multiplication query for slice %+v\n"+
			"Expected result: %d\n"+
			"Got instead: %d\n", s.GetArray(), 42, q)
	}

	s.Update(3, 2)
	q = s.Query(0, 2)
	if s.GetArray()[2] != 3 {
		t.Fatalf("Failed to update an element of the slice")
	}
	if q != 18 {
		t.Fatalf("Failed to calculate multiplication query for slice %+v after updating an element\n"+
			"Expected result: %d\n"+
			"Got instead: %d\n", s.GetArray(), 18, q)
	}

	s.UpdateRange([]int{3, 3, 3}, 3, 5)
	if !reflect.DeepEqual(s.GetArray(), []int{2, 3, 3, 3, 3, 3}) {
		t.Fatalf("Failed to update subslice"+
			"Expected %+v\n"+
			"Got instead %+v\n", []int{2, 3, 3, 3, 3, 3}, s.GetArray())
	}
	q = s.Query(0, 5)
	if q != 486 {
		t.Fatalf("Failed to calculate multiplication query for slice %+v after updating a subslice\n"+
			"Expected result: %d\n"+
			"Got instead: %d\n", s.GetArray(), 486, q)
	}
}

func TestStringConcatSegTree(t *testing.T) {
	var s SegTree[string]
	concatFunc := func(values ...string) string {
		c := ""
		for _, v := range values {
			c += v
		}

		return c
	}

	s.Setup([]string{"abc", "def", "xyz", "uuu", "fff"}, concatFunc)
	queries := [][]int{{0, 1}, {1, 3}, {2, 4}}
	expectedResults := []string{"abcdef", "defxyzuuu", "xyzuuufff"}
	for i := range queries {
		start := queries[i][0]
		end := queries[i][1]
		res := s.Query(start, end)
		if res != expectedResults[i] {
			errMsg := fmt.Sprintf("Calculated min query for elements from %d to %d"+
				"(inclusively) of an array %+v\n"+
				"Expected result: `%s`\n"+
				"Got instead: `%s`\n", start, end, s.GetArray(), expectedResults[i], res)

			t.Fatalf(errMsg)
		}
	}
}

type Coord struct {
	x int
	y int
}

func (c *Coord) Dist() int {
	// Chebyshev distance
	a := c.x
	if a < 0 {
		a = -a
	}
	b := c.y
	if b < 0 {
		b = -b
	}
	res := a
	if b > a {
		res = b
	}

	return res
}

func TestStructNodes(t *testing.T) {

	var s SegTree[Coord]
	maxRadiusVector := func(values ...Coord) Coord {
		var longestI int = 0
		var longest = values[0].Dist()
		for i := 1; i < len(values); i++ {
			if values[i].Dist() > longest {
				longestI = i
				longest = values[i].Dist()
			}
		}

		return values[longestI]
	}

	s.Setup([]Coord{Coord{1, 1}, Coord{2, 4}, Coord{-1, -1}, Coord{5, 3}}, maxRadiusVector)
	queries := [][]int{{0, 3}, {1, 2}}
	expectedResults := []Coord{Coord{5, 3}, Coord{2, 4}}
	for i := range queries {
		start := queries[i][0]
		end := queries[i][1]
		res := s.Query(start, end)
		if res != expectedResults[i] {
			errMsg := fmt.Sprintf("Calculated sum query for elements from %d to %d"+
				"(inclusively) of an array %+v\n"+
				"Expected result: %+v\n"+
				"Got instead: %+v\n", start, end, s.GetArray(), expectedResults[i], res)

			t.Fatalf(errMsg)
		}

	}
}
