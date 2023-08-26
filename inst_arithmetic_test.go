package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	i := Add{Imm(1)}

	assert.Nil(t, i.Validate())
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "add 1", i.String())

	i = Add{Reg(Acc)}
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "add acc", i.String())
}

func TestSub(t *testing.T) {
	i := Sub{Imm(1)}

	assert.Nil(t, i.Validate())
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "sub 1", i.String())

	i = Sub{Reg(Acc)}
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "sub acc", i.String())
}

func TestMul(t *testing.T) {
	i := Mul{Imm(1)}

	assert.Nil(t, i.Validate())
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "mul 1", i.String())

	i = Mul{Reg(Acc)}
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "mul acc", i.String())
}

func TestNot(t *testing.T) {
	i := Not{}

	assert.Nil(t, i.Validate())
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "not", i.String())
}
