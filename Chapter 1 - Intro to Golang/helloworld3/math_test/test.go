package math_test

import (
	"github.com/magiconair/properties/assert"
	"github.com/myles-mcdonnell/helloworld2/math"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, int64(5), math.Add(3,2))
}
