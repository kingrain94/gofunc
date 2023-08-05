package gofunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type T struct {
	X int64
	Y string
	Z string
}

func Test_Slice_Sorting(t *testing.T) {
	assert.Equal(t, []int{-5, -1, 0, 3}, Sort([]int{-1, 3, 0, -5}))
	assert.Equal(t, []string{"-5", "0", "5"}, Sort([]string{"0", "5", "-5"}))
	assert.Equal(t, []string{"aa", "ab", "j", "z"}, Sort([]string{"z", "j", "aa", "ab"}))
}

func Test_Slice_Sorting_Predicate(t *testing.T) {
	s := []T{
		{
			X: 10,
			Y: "aa",
			Z: "ab",
		},
		{
			X: -5,
			Y: "j",
			Z: "ac",
		},
		{
			X: 20,
			Y: "ab",
			Z: "z",
		},
	}
	assert.Equal(t, []T{
		{
			X: -5,
			Y: "j",
			Z: "ac",
		},
		{
			X: 10,
			Y: "aa",
			Z: "ab",
		},
		{
			X: 20,
			Y: "ab",
			Z: "z",
		},
	}, SortPred(s, func(item T) int64 {
		return item.X
	}))
	assert.Equal(t, []T{
		{
			X: 10,
			Y: "aa",
			Z: "ab",
		},
		{
			X: 20,
			Y: "ab",
			Z: "z",
		},
		{
			X: -5,
			Y: "j",
			Z: "ac",
		},
	}, SortPred(s, func(item T) string {
		return item.Y
	}))
	assert.Equal(t, []T{
		{
			X: 10,
			Y: "aa",
			Z: "ab",
		},
		{
			X: -5,
			Y: "j",
			Z: "ac",
		},
		{
			X: 20,
			Y: "ab",
			Z: "z",
		},
	}, SortPred(s, func(item T) string {
		return item.Z
	}))
}
