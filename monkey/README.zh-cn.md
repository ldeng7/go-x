概要
====

### e.g. 1 - 用函数替换函数
```
import (
	"github.com/ldeng7/go-x/monkey"
)

//go:noinline
func add1(i int) int {
	return i + 1
}

func add2(i int) int {
	return i + 2
}

func add3(i int) int {
	return i + 3
}

func main() {
	p, _ := monkey.NewPatch(add1, add2)
	p.Patch()
	println(add1(0)) // 2
	p.Unpatch()
	println(add1(0)) // 1
	p.Patch()
	println(add1(0)) // 2

	p1, _ := monkey.NewPatch(add1, add3)
	p1.Patch()
	println(add1(0)) // 3
	p1.Unpatch()
	println(add1(0)) // 2

	p.Unpatch()
	p, _ = monkey.NewPatch(add1, add2)
	p.Patch()
	p, _ = monkey.NewPatch(add2, add3)
	p.Patch()
	println(add1(0)) // 3
}
```

### e.g. 2 - 用函数替换方法
```
type S struct {
	i int
}

//go:noinline
func (s S) Add(i int) int {
	return s.i + i
}

func add1(i int) int {
	return i + 1
}

func sub(s S, i int) int {
	return s.i - i
}

func main() {
	s := &S{0}

	sAdd := s.Add
	p, _ := monkey.NewPatch(sAdd, add1)
	p.Patch()
	println(sAdd(1)) // 2

	p, _ = monkey.NewPatch(S.Add, sub)
	p.Patch()
	println(s.Add(1)) // -1
}
```

### e.g. 3 - 用方法替换方法
```
type S struct {
	i int
}

//go:noinline
func (s S) Add(i int) int {
	return s.i + i
}

func (s S) Sub(i int) int {
	return s.i - i
}

func main() {
	s := &S{0}
	p, _ := monkey.NewPatch(S.Add, S.Sub)
	p.Patch()
	println(s.Add(1)) // -1
}
```

### e.g. 4 - 你没有看错还有用方法替换函数的骚操作
```
//go:noinline
func add1(i int) int {
	return i + 1
}

type S struct {
	i int
}

func (s S) Add(i int) int {
	return s.i + i
}

func main() {
	s := &S{0}
	sAdd := s.Add
	p, _ := monkey.NewPatch(add1, sAdd)
	p.Patch()
	println(add1(1)) // 1
}
```

### e.g. 5 - 更骚的是在替代函数中还能调用原函数
```
//go:noinline
func add1(i int) int {
	return i + 1
}

func main() {
	var p *monkey.Patch
	add2 := func(i int) int {
		p.Unpatch()
		defer p.Patch()
		return add1(i) + 1
	}
	p, _ = monkey.NewPatch(add1, add2)
	p.Patch()
	println(add1(0)) // 2
}
```
