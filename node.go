package segtree

type node[T any] struct {
	value  T
	lBound int
	rBound int
	left   *node[T]
	right  *node[T]
}

func (n *node[T]) setup(arr []T, refArr []*T, f func(values ...T) T) {
	left := 0
	right := len(arr) - 1

	n.lBound = left
	n.rBound = right

	var set func(nod *node[T], l int, r int) T
	set = func(nod *node[T], l int, r int) T {
		nod.lBound = l
		nod.rBound = r

		if l == r {
			nod.value = arr[l]
			refArr[l] = &nod.value
			return nod.value
		}
		if l+1 == r {
			nod.value = f(arr[l], arr[r])
			nod.left = &node[T]{value: arr[l], lBound: l, rBound: l}
			refArr[l] = &nod.left.value
			nod.right = &node[T]{value: arr[r], lBound: r, rBound: r}
			refArr[r] = &nod.right.value
			return nod.value
		}
		mid := (l + r) / 2
		nod.left = &node[T]{}
		nod.right = &node[T]{}
		lValue := set(nod.left, l, mid)
		rValue := set(nod.right, mid+1, r)
		nod.value = f(lValue, rValue)

		return nod.value
	}

	n.value = set(n, left, right)
}

func (n *node[T]) update(val T, index int, f func(values ...T) T) {
	var crawler func(nod *node[T]) T
	crawler = func(nod *node[T]) T {
		if nod.lBound == index && nod.rBound == index {
			nod.value = val
			return nod.value
		}

		mid := (nod.lBound + nod.rBound) / 2
		if index <= mid {
			nod.value = f(crawler(nod.left), nod.right.value)
		} else {
			nod.value = f(crawler(nod.right), nod.left.value)
		}

		return nod.value
	}
	n.value = crawler(n)
}

func (n *node[T]) query(left int, right int, f func(values ...T) T) T {
	var records []T

	var q func(nod *node[T])

	q = func(nod *node[T]) {
		if left <= nod.lBound && nod.rBound <= right {
			records = append(records, nod.value)
			return
		}

		if left <= nod.left.rBound {
			q(nod.left)
		}

		if nod.right.lBound <= right {
			q(nod.right)
		}
	}
	q(n)

	sum := f(records...)

	return sum
}

func (n *node[T]) insert(arr []T, start int, end int, f func(values ...T) T) {
	var ins func(nod *node[T], l int, r int) T

	ins = func(nod *node[T], l int, r int) T {
		if nod.rBound < start || end < nod.lBound {
			return nod.value
		}

		if nod.lBound == nod.rBound {
			if start <= nod.lBound && nod.lBound <= end {
				nod.value = arr[nod.lBound-start]
				return nod.value
			}
		}

		nod.value = f(ins(nod.left, nod.left.lBound, nod.left.rBound),
			ins(nod.right, nod.right.lBound, nod.right.rBound))

		return nod.value
	}

	n.value = ins(n, start, end)
}
