package fengzhouemu

import (
	"fmt"
	"math"
)

type Add struct {
	A Operand
}

func (a Add) Validate() error {
	if imm, ok := a.A.(Imm); ok {
		return imm.Validate()
	}

	return nil
}

func (a Add) Cost() int {
	return 1
}

func (a Add) Execute(mc *MC) {
	v := a.A.Value(mc.reg)

	n := mc.reg[Acc].Read() + v
	mc.reg[Acc].Write(n)
}

func (a Add) Accesses() []Reg {
	if reg, ok := a.A.(Reg); ok && reg != Acc {
		return []Reg{Acc, reg}
	}

	return []Reg{Acc}
}

func (a Add) String() string {
	return fmt.Sprintf("add %v", a.A)
}

func (a Add) Label() string {
	return ""
}

func (a Add) Condition() ConditionType {
	return Always
}

type Sub struct {
	A Operand
}

func (s Sub) Validate() error {
	if imm, ok := s.A.(Imm); ok {
		return imm.Validate()
	}

	return nil
}

func (s Sub) Cost() int {
	return 1
}

func (s Sub) Execute(mc *MC) {
	v := s.A.Value(mc.reg)

	n := mc.reg[Acc].Read() - v
	mc.reg[Acc].Write(n)
}

func (s Sub) Accesses() []Reg {
	if reg, ok := s.A.(Reg); ok && reg != Acc {
		return []Reg{Acc, reg}
	}

	return []Reg{Acc}
}

func (s Sub) String() string {
	return fmt.Sprintf("sub %v", s.A)
}

func (s Sub) Label() string {
	return ""
}

func (s Sub) Condition() ConditionType {
	return Always
}

type Mul struct {
	A Operand
}

func (m Mul) Validate() error {
	if imm, ok := m.A.(Imm); ok {
		return imm.Validate()
	}

	return nil
}

func (m Mul) Cost() int {
	return 1
}

func (m Mul) Execute(mc *MC) {
	v := m.A.Value(mc.reg)

	n := mc.reg[Acc].Read() * v
	mc.reg[Acc].Write(n)
}

func (m Mul) Accesses() []Reg {
	if reg, ok := m.A.(Reg); ok && reg != Acc {
		return []Reg{Acc, reg}
	}

	return []Reg{Acc}
}

func (m Mul) String() string {
	return fmt.Sprintf("mul %v", m.A)
}

func (m Mul) Label() string {
	return ""
}

func (m Mul) Condition() ConditionType {
	return Always
}

type Not struct{}

func (n Not) Validate() error {
	return nil
}

func (n Not) Cost() int {
	return 1
}

func (n Not) Execute(mc *MC) {
	v := mc.reg[Acc].Read()

	switch v {
	case 0:
		v = 100
	default:
		v = 0
	}

	mc.reg[Acc].Write(v)
}

func (n Not) Accesses() []Reg {
	return []Reg{Acc}
}

func (n Not) String() string {
	return "not"
}

func (n Not) Label() string {
	return ""
}

func (n Not) Condition() ConditionType {
	return Always
}

type Dgt struct {
	A Operand
}

func (d Dgt) Validate() error {
	if imm, ok := d.A.(Imm); ok {
		return imm.Validate()
	}

	return nil
}

func (d Dgt) Cost() int {
	return 1
}

func (d Dgt) Execute(mc *MC) {
	digit := int(d.A.Value(mc.reg))
	if digit > 2 || digit < 0 {
		mc.reg[Acc].Write(0)
		return
	}

	acc := int(mc.reg[Acc].Read())
	m := int(math.Pow10(digit))

	mc.reg[Acc].Write(int16(acc/m) % 10)
}

func (d Dgt) Accesses() []Reg {
	if reg, ok := d.A.(Reg); ok && reg != Acc {
		return []Reg{Acc, reg}
	}

	return []Reg{Acc}
}

func (d Dgt) String() string {
	return fmt.Sprintf("dgt %v", d.A)
}

func (d Dgt) Label() string {
	return ""
}

func (d Dgt) Condition() ConditionType {
	return Always
}

type Dst struct {
	A Operand
	B Operand
}

func (d Dst) Validate() error {
	if imm, ok := d.A.(Imm); ok {
		if err := imm.Validate(); err != nil {
			return err
		}
	}
	if imm, ok := d.B.(Imm); ok {
		if err := imm.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (d Dst) Cost() int {
	return 1
}

func (d Dst) Execute(mc *MC) {
	digit := int(d.A.Value(mc.reg))
	if digit > 2 || digit < 0 {
		return
	}

	acc := int(mc.reg[Acc].Read())
	val := int(d.B.Value(mc.reg))
	m := int(math.Pow10(digit))

	// https://stackoverflow.com/a/3858110
	pre := acc / (m * 10) * (m * 10)
	suf := acc % m
	mid := val * m

	mc.reg[Acc].Write(int16(pre + mid + suf))
}

func (d Dst) Accesses() (registers []Reg) {
	registers = append(registers, Acc)
	if reg, ok := d.A.(Reg); ok && reg != Acc {
		registers = append(registers, reg)
	}
	if reg, ok := d.B.(Reg); ok && reg != Acc && reg != d.A {
		registers = append(registers, reg)
	}
	return
}

func (d Dst) String() string {
	return fmt.Sprintf("dst %v %v", d.A, d.B)
}

func (d Dst) Label() string {
	return ""
}

func (d Dst) Condition() ConditionType {
	return Always
}
