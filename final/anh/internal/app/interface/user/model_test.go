package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumber_IsMSISDN(t *testing.T) {
	assert.True(t, Number("13336061916").IsMSISDN())
	assert.False(t, Number("10000000000").IsMSISDN())
}
