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

func TestDgt(t *testing.T) {
	i := Dgt{Imm(0)}

	assert.Nil(t, i.Validate())
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "dgt 0", i.String())

	i = Dgt{Reg(Acc)}
	assert.Equal(t, []Reg{Acc}, i.Accesses())

	i = Dgt{Reg(Dat)}
	assert.Equal(t, []Reg{Acc, Dat}, i.Accesses())
}

func TestDgtExecute(t *testing.T) {
	vector := []struct {
		A      Operand
		Acc    int16
		Result int16
	}{
		{
			Imm(0), 596, 6,
		},
		{
			Imm(1), 596, 9,
		},
		{
			Imm(2), 596, 5,
		},
		{
			Imm(3), 596, 0,
		},
		{
			Imm(-1), 596, 0,
		},
	}

	for _, v := range vector {
		reg := newRegisters()
		reg[Acc].Write(v.Acc)

		i := Dgt{v.A}
		m := NewMC(reg)
		err := m.Load([]Inst{i})
		assert.Nil(t, err)

		i.Execute(m)
		assert.Equal(t, v.Result, reg[Acc].Read(), "A: %v, Acc: %v, Actual: %v", v.A, v.Acc, reg[Acc].Read())
	}
}

func TestDst(t *testing.T) {
	i := Dst{Imm(0), Imm(7)}

	assert.Nil(t, i.Validate())
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "dst 0 7", i.String())

	i = Dst{Reg(Dat), Reg(Acc)}
	assert.Equal(t, []Reg{Acc, Dat}, i.Accesses())
}

func TestDstExecute(t *testing.T) {
	vector := []struct {
		A      Operand
		B      Operand
		Acc    int16
		Result int16
	}{
		{
			Imm(0), Imm(7), 596, 597,
		},
		{
			Imm(1), Imm(7), 596, 576,
		},
		{
			Imm(2), Imm(7), 596, 796,
		},
		{
			Imm(3), Imm(7), 596, 596,
		},
		{
			Imm(2), Reg(Acc), 7, 707,
		},
		{
			Imm(2), Reg(Acc), -7, -707,
		},
	}

	for _, v := range vector {
		reg := newRegisters()
		reg[Acc].Write(v.Acc)

		i := Dst{v.A, v.B}
		m := NewMC(reg)
		err := m.Load([]Inst{i})
		assert.Nil(t, err)

		i.Execute(m)
		assert.Equal(t, v.Result, reg[Acc].Read(), "A: %v, B: %v, Acc: %v, Actual: %v", v.A, v.B, v.Acc, reg[Acc].Read())
	}
}
