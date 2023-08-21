package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMCEmptyProgram(t *testing.T) {
	m, _ := NewMC(defaultMC4000Registers, []Inst{})

	m.Step()
	assert.Equal(t, 0, m.ip)
	m.Step()
	assert.Equal(t, 0, m.ip)
}

func TestMCNonEmptyProgram(t *testing.T) {
	m, _ := NewMC(defaultMC4000Registers, []Inst{Nop{}, Nop{}})

	m.Step()
	assert.Equal(t, 1, m.ip)
	m.Step()
	assert.Equal(t, 0, m.ip)
	m.Step()
	assert.Equal(t, 1, m.ip)
}

func TestMCArithmetic(t *testing.T) {
	m, _ := NewMC(defaultMC4000Registers, []Inst{
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

func TestMCPower(t *testing.T) {
	m, _ := NewMC(defaultMC4000Registers, []Inst{
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
