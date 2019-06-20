package ints

type omKeyType = int
type omElemType = int

type ArrayOrderedMap struct {
	m  map[omKeyType]omElemType
	oa OrderedArray
}

func (om *ArrayOrderedMap) Init(lessCb oaElemLessCb) *ArrayOrderedMap {
	om.m = map[omKeyType]omElemType{}
	om.oa.Init(nil, lessCb)
	return om
}

func (om *ArrayOrderedMap) Len() int {
	return len(om.m)
}

func (om *ArrayOrderedMap) Get(key omKeyType) (omElemType, bool) {
	v, ok := om.m[key]
	return v, ok
}

func (om *ArrayOrderedMap) GetAt(index int) (omKeyType, omElemType) {
	k := om.oa.Get(index)
	v, _ := om.m[k]
	return k, v
}

func (om *ArrayOrderedMap) LowerBound(key omKeyType) int {
	return om.oa.LowerBound(key)
}

func (om *ArrayOrderedMap) UpperBound(key omKeyType) int {
	return om.oa.LowerBound(key)
}

func (om *ArrayOrderedMap) Set(key omKeyType, value omElemType) {
	if _, ok := om.m[key]; !ok {
		om.oa.Add(key)
	}
	om.m[key] = value
}

func (om *ArrayOrderedMap) RemoveAt(index int) {
	delete(om.m, om.oa.Get(index))
	om.oa.RemoveAt(index)
}

func (om *ArrayOrderedMap) RemoveRange(indexBegin, indexEnd int) {
	for i := indexBegin; i < indexEnd; i++ {
		delete(om.m, om.oa.Get(i))
	}
	om.oa.RemoveRange(indexBegin, indexEnd)
}

func (om *ArrayOrderedMap) Remove(key omKeyType) {
	delete(om.m, key)
	om.oa.Remove(key)
}
