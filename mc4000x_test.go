package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMC4000XValidation(t *testing.T) {
	// MC4000X does have an X2 register
	m := NewMC4000X()
	err := m.Load(MC4000XProgram{
		Mov{Imm(1), Reg(X2)},
	})
	assert.Nil(t, err)
	// assert.ErrorIs(t, err, PinNotConnectedErr{Reg(X2)})

	// MC4000X does NOT have a P0 register
	m = NewMC4000X()
	err = m.Load(MC4000XProgram{
		Mov{Imm(1), Reg(P0)},
	})
	assert.ErrorIs(t, err, InvalidRegisterErr{Reg(P0)})
}
