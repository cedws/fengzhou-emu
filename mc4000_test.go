package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMC4000Validation(t *testing.T) {
	// MC4000 does have a P0 register
	_, err := NewMC4000(MC4000Program{
		Mov{Imm(1), Reg(P0)},
	})
	assert.Nil(t, err)

	// MC4000 does NOT have an X2 register
	_, err = NewMC4000(MC4000Program{
		Mov{Imm(1), Reg(X2)},
	})
	assert.NotNil(t, err)
}
