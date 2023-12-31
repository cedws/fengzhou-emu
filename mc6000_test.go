package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMC6000Validation(t *testing.T) {
	// MC6000 does have an X2 register
	m := NewMC6000()
	err := m.Load(MC6000Program{
		Mov{Imm(1), Reg(X2)},
	})
	assert.Nil(t, err)
	// assert.ErrorIs(t, err, PinNotConnectedErr{Reg(X2)})

	// MC6000 does have a P0 register
	m = NewMC6000()
	err = m.Load(MC6000Program{
		Mov{Imm(1), Reg(P0)},
	})
	assert.ErrorIs(t, err, PinNotConnectedErr{Reg(P0)})
}
