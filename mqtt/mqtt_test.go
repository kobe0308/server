package mqtt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMqtt(t *testing.T) {
	// todo add mqtt tests...
	// arrange
	var a = 1
	var b = 2
	// act
	var res = a + b
	// assert
	assert.Equal(t, 3, res, "computer error again")
}
