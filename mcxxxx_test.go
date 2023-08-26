package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newRegisters() map[Reg]Register {
	return map[Reg]Register{
		Null:  NullRegister{},
		Acc:   &InternalRegister{},
		Dat:   &InternalRegister{},
		P0:    &SimplePinRegister{},
		P1:    &SimplePinRegister{},
		X0:    &XbusPinRegister{},
		X1:    &XbusPinRegister{},
		X2:    &XbusPinRegister{},
		X3:    &XbusPinRegister{},
		ip:    &InternalRegister{},
		flags: &InternalRegister{},
	}
}

func TestImmediate(t *testing.T) {
	assert.ErrorIs(t, Imm(-1000).Validate(), NumberTooSmallErr{-1000})
	assert.ErrorIs(t, Imm(1000).Validate(), NumberTooLargeErr{1000})
}

func TestMCEmptyProgram(t *testing.T) {
	m, err := NewMC(newRegisters(), []Inst{})
	assert.Nil(t, err)

	m.Step()
	assert.Equal(t, int16(0), m.reg[ip].Read())
	m.Step()
	assert.Equal(t, int16(0), m.reg[ip].Read())
}

func TestMCNonEmptyProgram(t *testing.T) {
	m, err := NewMC(newRegisters(), []Inst{Nop{}, Nop{}})
	assert.Nil(t, err)

	assert.Equal(t, int16(0), m.reg[ip].Read())
	m.Step()
	assert.Equal(t, int16(1), m.reg[ip].Read())
	m.Step()
	assert.Equal(t, int16(0), m.reg[ip].Read())
	m.Step()
	assert.Equal(t, int16(1), m.reg[ip].Read())
}

func TestMCArithmetic(t *testing.T) {
	m, err := NewMC(newRegisters(), []Inst{
		Mov{Imm(1), Reg(Acc)},
		Add{Reg(Acc)},
		Sub{Imm(2)},
		Not{},
		Add{Imm(1)},
		Mul{Imm(42)},
		Not{},
		Sub{Imm(999)},
	})
	assert.Nil(t, err)

	m.Step()
	assert.Equal(t, int16(1), m.reg[Acc].Read())
	m.Step()
	assert.Equal(t, int16(2), m.reg[Acc].Read())
	m.Step()
	assert.Equal(t, int16(0), m.reg[Acc].Read())
	m.Step()
	assert.Equal(t, int16(100), m.reg[Acc].Read())
	m.Step()
	assert.Equal(t, int16(101), m.reg[Acc].Read())
	m.Step()
	assert.Equal(t, int16(999), m.reg[Acc].Read())
	m.Step()
	assert.Equal(t, int16(0), m.reg[Acc].Read())
	m.Step()
	assert.Equal(t, int16(-999), m.reg[Acc].Read())
}

func TestMCPower(t *testing.T) {
	m, err := NewMC(newRegisters(), []Inst{
		Mov{Imm(1), Reg(Acc)},
		Add{Reg(Acc)},
		Add{Reg(Acc)},
		Sub{Imm(2)},
		Sub{Imm(2)},
		Add{Imm(1)},
		Mul{Imm(42)},
	})
	assert.Nil(t, err)

	for i := 0; i < 42; i++ {
		m.Step()
	}

	assert.Equal(t, 42, m.Power())
}

func TestMCNullRegister(t *testing.T) {
	m, err := NewMC(newRegisters(), []Inst{
		Mov{Imm(100), Reg(Null)},
	})
	assert.Nil(t, err)

	m.Step()
	assert.Equal(t, int16(0), m.reg[Null].Read())
}

func TestMCInternalRegister(t *testing.T) {
	m, err := NewMC(newRegisters(), []Inst{
		Mov{Imm(100), Reg(Acc)},
	})
	assert.Nil(t, err)

	m.Step()
	assert.Equal(t, int16(100), m.reg[Acc].Read())
}

func TestMCExecuteOnce(t *testing.T) {
	m, err := NewMC(newRegisters(), []Inst{
		Condition(Once, Mov{Imm(0), Reg(Acc)}),
		Condition(Once, Mov{Imm(1), Reg(Acc)}),
		Mov{Imm(2), Reg(Acc)},
	})
	assert.Nil(t, err)

	m.Step()
	assert.Equal(t, int16(0), m.reg[Acc].Read())
	m.Step()
	assert.Equal(t, int16(1), m.reg[Acc].Read())
	m.Step()
	assert.Equal(t, int16(2), m.reg[Acc].Read())
	m.Step()
	assert.Equal(t, int16(2), m.reg[Acc].Read())
	m.Step()
	assert.Equal(t, int16(2), m.reg[Acc].Read())
}

func TestMCExecuteEnable(t *testing.T) {
	m, err := NewMC(newRegisters(), []Inst{
		Condition(Enable, Mov{Imm(1), Reg(Acc)}),
		Teq{Imm(1), Imm(1)},
		Condition(Enable, Mov{Imm(1), Reg(Acc)}),
	})
	assert.Nil(t, err)

	m.Step()
	assert.Equal(t, int16(0), m.reg[Acc].Read())
	m.Step()
	m.Step()
	assert.Equal(t, int16(1), m.reg[Acc].Read())
}

func TestMCExecuteDisable(t *testing.T) {
	m, err := NewMC(newRegisters(), []Inst{
		Condition(Disable, Mov{Imm(1), Reg(Acc)}),
		Teq{Imm(1), Imm(2)},
		Condition(Disable, Mov{Imm(1), Reg(Acc)}),
	})
	assert.Nil(t, err)

	m.Step()
	assert.Equal(t, int16(0), m.reg[Acc].Read())
	m.Step()
	m.Step()
	assert.Equal(t, int16(1), m.reg[Acc].Read())
}

func TestMCTcpExecuteCondition(t *testing.T) {
	m, err := NewMC(newRegisters(), []Inst{
		Tcp{Imm(2), Imm(1)},
		Condition(Enable, Mov{Imm(1), Reg(Dat)}),
		Condition(Disable, Mov{Imm(-1), Reg(Dat)}),
		Mov{Imm(0), Reg(Dat)},
		Tcp{Imm(1), Imm(1)},
		Condition(Enable, Mov{Imm(1), Reg(Dat)}),
		Condition(Disable, Mov{Imm(-1), Reg(Dat)}),
	})
	assert.Nil(t, err)

	m.Step()
	m.Step()
	assert.Equal(t, int16(1), m.reg[Dat].Read())
	m.Step()
	assert.Equal(t, int16(0), m.reg[Dat].Read())
	m.Step()
	assert.Equal(t, int16(0), m.reg[Dat].Read())
}

func TestMCLabelAlreadyDefined(t *testing.T) {
	_, err := NewMC(newRegisters(), []Inst{
		Label("123", Mov{Imm(100), Reg(Dat)}),
		Label("123", Mov{Imm(100), Reg(Dat)}),
	})
	assert.ErrorIs(t, err, LabelAlreadyDefinedErr{"123"})
}

func TestMCUndefinedJmp(t *testing.T) {
	_, err := NewMC(newRegisters(), []Inst{
		Jmp{"123"},
	})
	assert.ErrorIs(t, err, LabelNotDefinedErr{"123"})
}

func TestMCDefinedJmp(t *testing.T) {
	_, err := NewMC(newRegisters(), []Inst{
		Jmp{"123"},
		Label("123", Nop{}),
	})
	assert.Nil(t, err)
}

func TestMCJmp(t *testing.T) {
	m, err := NewMC(newRegisters(), []Inst{
		Jmp{"123"},
		Nop{},
		Nop{},
		Nop{},
		Label("123", Not{}),
	})
	assert.Nil(t, err)

	m.Step()
	m.Step()
	assert.Equal(t, int16(100), m.reg[Acc].Read())
	assert.Equal(t, int(2), m.Power())
	m.Step()
	m.Step()
	assert.Equal(t, int16(0), m.reg[Acc].Read())
	assert.Equal(t, int(4), m.Power())
}
