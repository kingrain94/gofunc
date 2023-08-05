package gofunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cond_If(t *testing.T) {
	type foo struct {
		Val int
	}
	assert.Equal(t, "first", If[interface{}](true, "first", 2))
	assert.Equal(t, "second", If(false, "first", "second"))
	assert.Equal(t, 2.2, If(false, -1.0, 2.2))
	assert.Equal(t, foo{Val: 1}, If(true, foo{Val: 1}, foo{Val: 2}))
	assert.Equal(t, []int{1, 2, 3}, If(true, []int{1, 2, 3}, []int{4, 5, 6}))
}
