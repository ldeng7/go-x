package structs

type ArrayUnion struct {
	arr []int
}

func (au *ArrayUnion) Init(l int) *ArrayUnion {
	au.arr = make([]int, l)
	for i := 0; i < l; i++ {
		au.arr[i] = -1
	}
	return au
}

func (au *ArrayUnion) Set(i, v int) {
	au.arr[i] = v
}

func (au *ArrayUnion) GetRoot(i int) int {
	r := au.arr[i]
	if r == -1 || r == i {
		return i
	}
	r = au.GetRoot(r)
	au.arr[i] = r
	return r
}

func (au *ArrayUnion) Merge(i, j int) bool {
	r1, r2 := au.GetRoot(i), au.GetRoot(j)
	if r1 != r2 {
		au.arr[r1] = r2
		return true
	}
	return false
}

func (au *ArrayUnion) Get(i int) int {
	return au.arr[i]
}

func (au *ArrayUnion) NumGroup() int {
	ret := 0
	for i, r := range au.arr {
		if r == -1 || r == i {
			ret++
		}
	}
	return ret
}

func (au *ArrayUnion) Reset() {
	l := len(au.arr)
	for i := 0; i < l; i++ {
		au.arr[i] = -1
	}
}

func (au *ArrayUnion) Copy() *ArrayUnion {
	arr := make([]int, len(au.arr))
	copy(arr, au.arr)
	return &ArrayUnion{arr}
}
