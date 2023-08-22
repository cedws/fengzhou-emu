package fengzhouemu

import "fmt"

const (
	regMin = -999
	regMax = 999
)

type ConditionType int

const (
	Always ConditionType = iota
	Once
	Enable
	Disable
)

type condition struct {
	Inst
	c ConditionType
}

type labelled struct {
	Inst
	l string
}

func (l labelled) Label() string {
	return l.l
}

func Label(l string, i Inst) Inst {
	return labelled{
		Inst: i,
		l:    l,
	}
}

func (c condition) Condition() ConditionType {
	return c.c
}

func Condition(c ConditionType, i Inst) Inst {
	return condition{
		Inst: i,
		c:    c,
	}
}

// Operand represents a generic operand that can be either a Register or immediate value.
type Operand interface {
	Value(map[Reg]Register) int16
}

// Inst represents a single instruction.
type Inst interface {
	Validate() error
	Execute(map[Reg]Register)
	Cost() int
	Accesses() []Reg
	String() string
	Label() string
	Condition() ConditionType
}

type Nop struct{}

func (n Nop) Validate() error {
	return nil
}

func (n Nop) Cost() int {
	return 1
}

func (n Nop) Execute(registers map[Reg]Register) {}

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

func (m Mov) Execute(registers map[Reg]Register) {
	v := m.A.Value(registers)
	registers[m.B].Write(v)
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

func (a Add) Execute(registers map[Reg]Register) {
	v := a.A.Value(registers)

	n := registers[Acc].Read() + v
	n = min(n, regMax)
	n = max(n, regMin)

	registers[Acc].Write(n)
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

func (s Sub) Execute(registers map[Reg]Register) {
	v := s.A.Value(registers)

	n := registers[Acc].Read() - v
	n = min(n, regMax)
	n = max(n, regMin)

	registers[Acc].Write(n)
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

func (m Mul) Execute(registers map[Reg]Register) {
	v := m.A.Value(registers)

	n := registers[Acc].Read() * v
	n = min(n, regMax)
	n = max(n, regMin)

	registers[Acc].Write(n)
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
