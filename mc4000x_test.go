package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMC4000XValidation(t *testing.T) {
	// MC4000X does have an X2 register
	_, err := NewMC4000X(MC4000XProgram{
		Mov{Imm(1), Reg(X2)},
	})
	assert.Nil(t, err)

	// MC4000X does NOT have a P0 register
	_, err = NewMC4000X(MC4000XProgram{
		Mov{Imm(1), Reg(P0)},
	})
	assert.NotNil(t, err)
}
