package fengzhouemu

import (
	"fmt"
)

type Nop struct{}

func (n Nop) Validate() error {
	return nil
}

func (n Nop) Cost() int {
	return 1
}

func (n Nop) Execute(mc *MC) {}

func (n Nop) Accesses() []Reg {
	return nil
}

func (n Nop) String() string {
	return "nop"
}

func (n Nop) Label() string {
	return ""
}

func (n Nop) Condition() ConditionType {
	return Always
}

type Mov struct {
	A Operand
	B Reg
}

func (m Mov) Validate() error {
	if imm, ok := m.A.(Imm); ok {
		return imm.Validate()
	}

	return nil
}

func (m Mov) Cost() int {
	return 1
}

func (m Mov) Execute(mc *MC) {
	v := m.A.Value(mc.reg)
	mc.reg[m.B].Write(v)
}

func (m Mov) Accesses() []Reg {
	if reg, ok := m.A.(Reg); ok && reg != m.B {
		return []Reg{reg, m.B}
	}

	return []Reg{m.B}
}

func (m Mov) String() string {
	return fmt.Sprintf("mov %v %v", m.A, m.B)
}

func (m Mov) Label() string {
	return ""
}

func (m Mov) Condition() ConditionType {
	return Always
}

type Jmp struct {
	A string
}

func (j Jmp) Validate() error {
	return nil
}

func (j Jmp) Cost() int {
	return 1
}

func (j Jmp) Execute(mc *MC) {
	mc.jump(j.A)
}

func (j Jmp) Accesses() []Reg {
	return nil
}

func (j Jmp) String() string {
	return fmt.Sprintf("jmp %v", j.A)
}

func (j Jmp) Label() string {
	return ""
}

func (j Jmp) Condition() ConditionType {
	return Always
}
