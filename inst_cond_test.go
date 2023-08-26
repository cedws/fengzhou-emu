package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeq(t *testing.T) {
	i := Teq{Reg(Acc), Reg(Acc)}
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "teq acc acc", i.String())
}

func TestTeqExecute(t *testing.T) {
	reg := newRegisters()

	reg[Acc].Write(1)
	reg[Dat].Write(1)

	i := Teq{Reg(Acc), Reg(Dat)}
	m, err := NewMC(reg, []Inst{i})
	assert.Nil(t, err)

	i.Execute(m)
	assert.Equal(t, int16(testFlag|enableFlag), reg[flags].Read())

	reg[Acc].Write(2)
	reg[Dat].Write(1)

	i.Execute(m)
	assert.Equal(t, int16(testFlag), reg[flags].Read())
}

func TestTgt(t *testing.T) {
	i := Tgt{Reg(Acc), Reg(Acc)}
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "tgt acc acc", i.String())
}

func TestTgtExecute(t *testing.T) {
	reg := newRegisters()

	reg[Acc].Write(2)
	reg[Dat].Write(1)

	i := Tgt{Reg(Acc), Reg(Dat)}
	m, err := NewMC(reg, []Inst{i})
	assert.Nil(t, err)

	i.Execute(m)
	assert.Equal(t, int16(testFlag|enableFlag), reg[flags].Read())

	reg[Acc].Write(1)
	reg[Dat].Write(1)

	i.Execute(m)
	assert.Equal(t, int16(testFlag), reg[flags].Read())
}

func TestTlt(t *testing.T) {
	i := Tlt{Reg(Acc), Reg(Acc)}
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "tlt acc acc", i.String())
}

func TestTltExecute(t *testing.T) {
	reg := newRegisters()

	reg[Acc].Write(1)
	reg[Dat].Write(2)

	i := Tlt{Reg(Acc), Reg(Dat)}
	m, err := NewMC(reg, []Inst{i})
	assert.Nil(t, err)

	i.Execute(m)
	assert.Equal(t, int16(testFlag|enableFlag), reg[flags].Read())

	reg[Acc].Write(1)
	reg[Dat].Write(1)

	i.Execute(m)
	assert.Equal(t, int16(testFlag), reg[flags].Read())
}

func TestTcp(t *testing.T) {
	i := Tcp{Reg(Acc), Reg(Acc)}
	assert.Equal(t, []Reg{Acc}, i.Accesses())
	assert.Equal(t, "tcp acc acc", i.String())
}

func TestTcpExecute(t *testing.T) {
	reg := newRegisters()

	reg[Acc].Write(2)
	reg[Dat].Write(1)

	i := Tcp{Reg(Acc), Reg(Dat)}
	m, err := NewMC(reg, []Inst{i})
	assert.Nil(t, err)

	i.Execute(m)
	assert.Equal(t, int16(testFlag|enableFlag), reg[flags].Read())

	reg[Acc].Write(1)
	reg[Dat].Write(2)

	i.Execute(m)
	assert.Equal(t, int16(testFlag), reg[flags].Read())

	reg[Acc].Write(1)
	reg[Dat].Write(1)

	i.Execute(m)
	assert.Equal(t, int16(0), reg[flags].Read())
}
