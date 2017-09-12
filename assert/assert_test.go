package assert

import (
	"testing"
)

type ss struct {
	i int
	t tt
}

type tt struct {
	i int
}

func Test(t *testing.T) {
	ast := New(t)
	ast.True(true)
	ast.False(false)
	ast.Equal(1, 1)
	i1, i2 := 1, 1
	ast.Equal(&i1, &i2)
	ast.Equal("ab", "a"+"b")
	ast.Equal([]int{1, 2}, []int{1, 2})
	ast.Equal(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 2, "a": 1})
	ast.Equal(tt{i: 1}, tt{i: 1})
	ast.Equal([]tt{tt{i: 1}}, []tt{tt{i: 1}})
	ast.Equal(ss{i: 1, t: tt{i: 2}}, ss{i: 1, t: tt{i: 2}})
	ast.NotEqual(1, 0)
}
