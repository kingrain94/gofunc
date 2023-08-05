package gofunc

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Util_Must(t *testing.T) {
	// No error case
	assert.Equal(t, 1, Must(func() (int, error) { return 1, nil }()))
	assert.Equal(t, "a", Must(func() (string, error) { return "a", nil }()))

	// Error case (will panic)
	defer func() {
		e := recover()
		assert.True(t, e != nil && e.(error).Error() == "error")
	}()
	assert.Equal(t, 0, Must(func() (int, error) { return 0, errors.New("error") }()))
}
