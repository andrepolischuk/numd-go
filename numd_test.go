package numd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecline(t *testing.T) {
	assert.Equal(t, Decline(1, "рубль", "рубля", "рублей"), "1 рубль")
	assert.Equal(t, Decline(-4, "рубль", "рубля", "рублей"), "-4 рубля")
	assert.Equal(t, Decline(7, "рубль", "рубля", "рублей"), "7 рублей")
	assert.Equal(t, Decline(10, "рубль", "рубля", "рублей"), "10 рублей")
	assert.Equal(t, Decline(-11, "рубль", "рубля", "рублей"), "-11 рублей")
	assert.Equal(t, Decline(17, "рубль", "рубля", "рублей"), "17 рублей")
	assert.Equal(t, Decline(31, "рубль", "рубля", "рублей"), "31 рубль")
	assert.Equal(t, Decline(34, "рубль", "рубля", "рублей"), "34 рубля")
	assert.Equal(t, Decline(-37, "рубль", "рубля", "рублей"), "-37 рублей")
	assert.Equal(t, Decline(101, "рубль", "рубля", "рублей"), "101 рубль")
	assert.Equal(t, Decline(104, "рубль", "рубля", "рублей"), "104 рубля")
	assert.Equal(t, Decline(107, "рубль", "рубля", "рублей"), "107 рублей")
	assert.Equal(t, Decline(1, "dollar", "dollars"), "1 dollar")
	assert.Equal(t, Decline(-4, "dollar", "dollars"), "-4 dollars");
	assert.Equal(t, Decline(7, "dollar", "dollars"), "7 dollars");
	assert.Equal(t, Decline(10, "dollar", "dollars"), "10 dollars");
	assert.Equal(t, Decline(-11, "dollar", "dollars"), "-11 dollars");
	assert.Equal(t, Decline(17, "dollar", "dollars"), "17 dollars");
	assert.Equal(t, Decline(31, "dollar", "dollars"), "31 dollars");
	assert.Equal(t, Decline(34, "dollar", "dollars"), "34 dollars");
	assert.Equal(t, Decline(-37, "dollar", "dollars"), "-37 dollars");
	assert.Equal(t, Decline(101, "dollar", "dollars"), "101 dollars")
	assert.Equal(t, Decline(104, "dollar", "dollars"), "104 dollars");
	assert.Equal(t, Decline(107, "dollar", "dollars"), "107 dollars");
}
