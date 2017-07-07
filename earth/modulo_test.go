package earth_test

import (
	"testing"

	"github.com/regattebzh/trajectory/earth"
	"github.com/stretchr/testify/assert"
)

func TestModuloPositive(t *testing.T) {
	angle := earth.Modulo(3650, 360)
	assert.Equal(t, int(50), angle, "Should be equals")
}

func TestModuloNegative(t *testing.T) {
	angle := earth.Modulo(-3650, 360)
	assert.Equal(t, int(310), angle, "Should be equals")
}
