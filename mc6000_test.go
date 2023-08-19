package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMC6000Validation(t *testing.T) {
	// MC6000 does have an X2 register
	_, err := NewMC6000(MC6000Program{
		Mov{Imm(1), Reg(X2)},
	})
	assert.Nil(t, err)

	// MC6000 does have a P0 register
	_, err = NewMC6000(MC6000Program{
		Mov{Imm(1), Reg(P0)},
	})
	assert.Nil(t, err)
}
