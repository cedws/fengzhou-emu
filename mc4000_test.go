package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMC4000Validation(t *testing.T) {
	// MC4000 does have a P0 register
	m := NewMC4000()
	err := m.Load(MC4000Program{
		Mov{Imm(1), Reg(P0)},
	})
	assert.ErrorIs(t, err, PinNotConnectedErr{Reg(P0)})

	// MC4000 does NOT have an X2 register
	m = NewMC4000()
	err = m.Load(MC4000Program{
		Mov{Imm(1), Reg(X2)},
	})
	assert.ErrorIs(t, err, InvalidRegisterErr{Reg(X2)})
}
