package funcx

type variant struct {
	index uint8
	data  any
}

func (v *variant) Index() uint8 {
	return v.index
}

type Variant2[T0, T1 any] struct {
	variant
}

func NewVariant2A[T0, T1 any](a T0) *Variant2[T0, T1] {
	return &Variant2[T0, T1]{variant{0, a}}
}

func NewVariant2B[T0, T1 any](b T1) *Variant2[T0, T1] {
	return &Variant2[T0, T1]{variant{0, b}}
}

func (v *Variant2[T0, T1]) A() (T0, bool) {
	a, ok := v.data.(T0)
	return a, ok
}

func (v *Variant2[T0, T1]) B() (T1, bool) {
	b, ok := v.data.(T1)
	return b, ok
}

func (v *Variant2[T0, T1]) SetA(a T0) {
	if v.index == 0 {
		v.data = a
	}
}

func (v *Variant2[T0, T1]) SetB(b T1) {
	if v.index == 1 {
		v.data = b
	}
}

type Variant3[T0, T1, T2 any] struct {
	variant
}

func NewVariant3A[T0, T1, T2 any](a T0) *Variant3[T0, T1, T2] {
	return &Variant3[T0, T1, T2]{variant{0, a}}
}

func NewVariant3B[T0, T1, T2 any](b T1) *Variant3[T0, T1, T2] {
	return &Variant3[T0, T1, T2]{variant{0, b}}
}

func NewVariant3C[T0, T1, T2 any](c T2) *Variant3[T0, T1, T2] {
	return &Variant3[T0, T1, T2]{variant{0, c}}
}

func (v *Variant3[T0, T1, T2]) A() (T0, bool) {
	a, ok := v.data.(T0)
	return a, ok
}

func (v *Variant3[T0, T1, T2]) B() (T1, bool) {
	b, ok := v.data.(T1)
	return b, ok
}

func (v *Variant3[T0, T1, T2]) C() (T2, bool) {
	c, ok := v.data.(T2)
	return c, ok
}

func (v *Variant3[T0, T1, T2]) SetA(a T0) {
	if v.index == 0 {
		v.data = a
	}
}

func (v *Variant3[T0, T1, T2]) SetB(b T1) {
	if v.index == 1 {
		v.data = b
	}
}

func (v *Variant3[T0, T1, T2]) SetC(c T2) {
	if v.index == 2 {
		v.data = c
	}
}
