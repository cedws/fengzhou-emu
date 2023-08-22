package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImmediate(t *testing.T) {
	i := Mov{Imm(-1000), Reg(Acc)}
	assert.ErrorIs(t, NumberTooSmallErr{-1000}, i.Validate())

	i = Mov{Imm(1000), Reg(Acc)}
	assert.ErrorIs(t, NumberTooLargeErr{1000}, i.Validate())
}

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

func TestInstLabelAndCondition(t *testing.T) {
	n := Nop{}
	assert.Equal(t, "", n.Label())

	nl := Label("start", n)
	assert.Equal(t, "start", nl.Label())

	assert.Equal(t, Always, n.Condition())
	nl = Condition(Disable, n)
	assert.Equal(t, Disable, nl.Condition())
}

func TestInstInvalidLabel(t *testing.T) {
	n := Label("hello;", Nop{})
	assert.Equal(t, "hello;", n.Label())

	_, err := NewMC4000(MC4000Program{n})
	assert.ErrorIs(t, InvalidLabelNameErr{"hello;"}, err)
}
