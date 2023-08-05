package gofunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Min(t *testing.T) {
	assert.Equal(t, 1, Must(Min(1)))
	assert.Equal(t, -10, Must(Min(0, 2, -10, -5, 3, 5)))
	assert.Equal(t, float32(-0.2), Must(Min[float32](0.1, -0.2, 0, 0, -0.2, 10)))
	assert.Equal(t, "01", Must(Min("1", "01", "10")))

	// Error case
	_, e := Min[int]()
	assert.ErrorIs(t, e, ErrInputRequired)
}

func Test_Max(t *testing.T) {
	assert.Equal(t, 1, Must(Max(1)))
	assert.Equal(t, 30, Must(Max(0, 2, -10, -5, 30, 5, 30)))
	assert.Equal(t, 10.11, Must(Max[float64](0.1, -0.2, 10.1, 0, -0.2, 10, 10.11)))
	assert.Equal(t, "10", Must(Max("1", "01", "10")))

	// Error case
	_, e := Max[string]()
	assert.ErrorIs(t, e, ErrInputRequired)
}
