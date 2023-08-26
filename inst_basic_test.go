package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNop(t *testing.T) {
	i := Nop{}

	assert.Nil(t, i.Validate())
}

func TestMov(t *testing.T) {
	i := Mov{Imm(1), Reg(Acc)}

	assert.Nil(t, i.Validate())
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "mov 1 acc", i.String())

	i = Mov{Reg(Acc), Reg(Acc)}
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "mov acc acc", i.String())
}
