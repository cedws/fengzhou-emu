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
		Sub{Imm(2)},
		Not{},
		Add{Imm(1)},
		Mul{Imm(42)},
		Not{},
	})

	m.Step()
	assert.Equal(t, int16(1), m.registers[Acc].Read())
	m.Step()
	assert.Equal(t, int16(2), m.registers[Acc].Read())
	m.Step()
	assert.Equal(t, int16(0), m.registers[Acc].Read())
	m.Step()
	assert.Equal(t, int16(100), m.registers[Acc].Read())
	m.Step()
	assert.Equal(t, int16(101), m.registers[Acc].Read())
	m.Step()
	assert.Equal(t, int16(999), m.registers[Acc].Read())
	m.Step()
	assert.Equal(t, int16(0), m.registers[Acc].Read())
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

func TestMCNullRegister(t *testing.T) {
	m, _ := NewMC(defaultMC4000Registers, []Inst{
		Mov{Imm(100), Reg(Null)},
	})

	m.Step()
	assert.Equal(t, int16(0), m.registers[Null].Read())
}

func TestMCInternalRegister(t *testing.T) {
	m, _ := NewMC(defaultMC4000Registers, []Inst{
		Mov{Imm(100), Reg(Acc)},
	})

	m.Step()
	assert.Equal(t, int16(100), m.registers[Acc].Read())
}
