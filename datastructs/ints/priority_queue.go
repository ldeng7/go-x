package ints

type PriorityQueue struct {
	Arr    []int
	lessCb func(a, b int) bool
}

func (pq *PriorityQueue) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !pq.lessCb(pq.Arr[j], pq.Arr[i]) {
			break
		}
		pq.Arr[i], pq.Arr[j] = pq.Arr[j], pq.Arr[i]
		j = i
	}
}

func (pq *PriorityQueue) down(i0, n int) bool {
	i := i0
	for {
		j1 := i<<1 + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && pq.lessCb(pq.Arr[j2], pq.Arr[j1]) {
			j = j2
		}
		if !pq.lessCb(pq.Arr[j], pq.Arr[i]) {
			break
		}
		pq.Arr[i], pq.Arr[j] = pq.Arr[j], pq.Arr[i]
		i = j
	}
	return i > i0
}

func (pq *PriorityQueue) Init(arr []int, lessCb func(int, int) bool) *PriorityQueue {
	pq.Arr = arr
	pq.lessCb = lessCb
	if nil == lessCb {
		pq.lessCb = func(a, b int) bool { return a < b }
	}
	l := len(pq.Arr)
	for i := l>>1 - 1; i >= 0; i-- {
		pq.down(i, l)
	}
	return pq
}

func (pq *PriorityQueue) Push(item int) {
	pq.Arr = append(pq.Arr, item)
	pq.up(len(pq.Arr) - 1)
}

func (pq *PriorityQueue) Pop() (int, bool) {
	i := len(pq.Arr) - 1
	if i >= 0 {
		pq.Arr[0], pq.Arr[i] = pq.Arr[i], pq.Arr[0]
		pq.down(0, i)
		v := pq.Arr[i]
		pq.Arr = pq.Arr[:i]
		return v, true
	}
	return 0, false
}

func (pq *PriorityQueue) RemoveAt(index int) {
	j := len(pq.Arr) - 1
	if j != index {
		pq.Arr[index], pq.Arr[j] = pq.Arr[j], pq.Arr[index]
		if !pq.down(index, j) {
			pq.up(index)
		}
	}
	pq.Arr = pq.Arr[:j]
}

func (pq *PriorityQueue) Set(index, item int) {
	pq.Arr[index] = item
	if !pq.down(index, len(pq.Arr)) {
		pq.up(index)
	}
}

func (pq *PriorityQueue) InitAsReversed(arr []int) *PriorityQueue {
	return pq.Init(arr, func(a, b int) bool { return a > b })
}
