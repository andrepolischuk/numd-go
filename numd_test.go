package numd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecline(t *testing.T) {
	assert.Equal(t, Decline(1, "рубль", "рубля", "рублей"), "1 рубль")
	assert.Equal(t, Decline(4, "доллар", "доллара", "долларов"), "4 доллара")
	assert.Equal(t, Decline(7, "франк", "франка", "франков"), "7 франков")
}
