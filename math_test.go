package gofunc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Abs(t *testing.T) {
	assert.Equal(t, int64(1), Abs(1))
	assert.Equal(t, int64(10), Abs(-10))
	assert.Equal(t, int64(0), Abs(0))

	// Error case
	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("function should panic")
			}
		}()

		Abs(math.MinInt64)
	})
}
