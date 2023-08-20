package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMC4000EmptyProgram(t *testing.T) {
	m, _ := NewMC4000(MC4000Program{})

	m.Step()
	assert.Equal(t, byte(0), m.ip)
	m.Step()
	assert.Equal(t, byte(0), m.ip)
	m.Step()
	assert.Equal(t, byte(0), m.ip)
}

func TestMC4000Arithmetic(t *testing.T) {
	m, _ := NewMC4000(MC4000Program{
		Mov{Imm(1), Reg(Acc)},
		Add{Reg(Acc)},
		Add{Reg(Acc)},
		Sub{Imm(2)},
		Sub{Imm(2)},
		Add{Imm(1)},
		Mul{Imm(42)},
	})

	m.Step()
	assert.Equal(t, int16(1), m.registers[Acc])
	m.Step()
	assert.Equal(t, int16(2), m.registers[Acc])
	m.Step()
	assert.Equal(t, int16(4), m.registers[Acc])
	m.Step()
	assert.Equal(t, int16(2), m.registers[Acc])
	m.Step()
	assert.Equal(t, int16(0), m.registers[Acc])
	m.Step()
	assert.Equal(t, int16(1), m.registers[Acc])
	m.Step()
	assert.Equal(t, int16(42), m.registers[Acc])
	m.Step()
	assert.Equal(t, int16(1), m.registers[Acc])
}

func TestMC4000Power(t *testing.T) {
	m, _ := NewMC4000(MC4000Program{
		Mov{Imm(1), Reg(Acc)},
		Add{Reg(Acc)},
		Add{Reg(Acc)},
		Sub{Imm(2)},
		Sub{Imm(2)},
		Add{Imm(1)},
		Mul{Imm(42)},
	})

	for i := 0; i < 42; i++ {
		m.Step()
	}

	assert.Equal(t, 42, m.Power())
}

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
