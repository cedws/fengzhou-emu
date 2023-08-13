package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArithmetic(t *testing.T) {
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
}
