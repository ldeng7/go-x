package assert

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/kr/pretty"
)

type Assert struct {
	t *testing.T
}

func New(t *testing.T) *Assert {
	return &Assert{t}
}

func (a *Assert) assert(b bool, fn func(), cd int) {
	if b {
		return
	}
	_, file, line, _ := runtime.Caller(cd + 1)
	a.t.Errorf("%s:%d", file, line)
	if nil != fn {
		fn()
	}
	a.t.FailNow()
}

func (a *Assert) Equal(got, exp interface{}) {
	fn := func() {
		for _, desc := range pretty.Diff(got, exp) {
			a.t.Error("!", desc)
		}
	}
	a.assert(reflect.DeepEqual(got, exp), fn, 1)
}

func (a *Assert) NotEqual(got, exp interface{}) {
	a.assert(!reflect.DeepEqual(got, exp), nil, 1)
}

func (a *Assert) True(b bool) {
	a.assert(b, nil, 1)
}

func (a *Assert) False(b bool) {
	a.assert(!b, nil, 1)
}

func (a *Assert) Fail(desc string) {
	fn := func() {
		a.t.Error("!", desc)
	}
	a.assert(false, fn, 1)
}
