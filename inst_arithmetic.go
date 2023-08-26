package fengzhouemu

import "fmt"

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
