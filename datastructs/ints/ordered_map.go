package ints

type OrderedMap struct {
	m  map[int]int
	oa OrderedArray
}

func (om *OrderedMap) Init(lessCb func(int, int) bool) *OrderedMap {
	om.m = map[int]int{}
	om.oa.Init(nil, lessCb)
	return om
}

func (om *OrderedMap) Len() int {
	return len(om.m)
}

func (om *OrderedMap) Get(k int) (int, bool) {
	v, ok := om.m[k]
	return v, ok
}

func (om *OrderedMap) Key(index int) int {
	return om.oa.Arr[index]
}

func (om *OrderedMap) Range(f func(k int, v int) bool) int {
	i := 0
	for _, k := range om.oa.Arr {
		if !f(k, om.m[k]) {
			break
		}
		i++
	}
	return i
}

func (om *OrderedMap) Set(k int, v int) {
	if _, ok := om.m[k]; !ok {
		om.oa.Add(k)
	}
	om.m[k] = v
}

func (om *OrderedMap) Inc(k int) {
	if _, ok := om.m[k]; !ok {
		om.oa.Add(k)
	}
	om.m[k]++
}

func (om *OrderedMap) Dec(k int) {
	if _, ok := om.m[k]; !ok {
		om.oa.Add(k)
	}
	om.m[k]--
}

func (om *OrderedMap) Remove(k int) {
	delete(om.m, k)
	om.oa.RemoveAt(om.oa.Index(k))
}
