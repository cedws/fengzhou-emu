package fengzhouemu

import "fmt"

const (
	regMin = -999
	regMax = 999
)

type (
	NumberTooSmallErr  struct{}
	NumberTooLargeErr  struct{}
	InvalidRegisterErr struct {
		Reg string
	}
)

func (e NumberTooSmallErr) Error() string {
	return "number too small"
}

func (e NumberTooLargeErr) Error() string {
	return "number too large"
}

func (e InvalidRegisterErr) Error() string {
	return fmt.Sprintf("invalid register %v", e.Reg)
}

type Operand interface {
	Value(map[Reg]int16) int16
}

type Inst interface {
	Validate() error
	Execute(map[Reg]int16)
	Cost() int
	Accesses() []Reg
	String() string
}

type Nop struct{}

func (n Nop) Validate() error {
	return nil
}

func (n Nop) Cost() int {
	return 1
}

func (n Nop) Execute(registers map[Reg]int16) {}

func (n Nop) Accesses() []Reg {
	return nil
}

func (n Nop) String() string {
	return "nop"
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

func (m Mov) Execute(registers map[Reg]int16) {
	v := m.A.Value(registers)
	registers[m.B] = v
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

func (a Add) Execute(registers map[Reg]int16) {
	v := a.A.Value(registers)
	registers[Acc] += v
	registers[Acc] = min(registers[Acc], regMax)
	registers[Acc] = max(registers[Acc], regMin)
}

func (a Add) Accesses() []Reg {
	if reg, ok := a.A.(Reg); ok {
		return []Reg{reg}
	}

	return nil
}

func (a Add) String() string {
	return fmt.Sprintf("add %v", a.A)
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

func (s Sub) Execute(registers map[Reg]int16) {
	v := s.A.Value(registers)
	registers[Acc] -= v
	registers[Acc] = min(registers[Acc], regMax)
	registers[Acc] = max(registers[Acc], regMin)
}

func (s Sub) Accesses() []Reg {
	if reg, ok := s.A.(Reg); ok {
		return []Reg{reg}
	}

	return nil
}

func (s Sub) String() string {
	return fmt.Sprintf("sub %v", s.A)
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

func (m Mul) Execute(registers map[Reg]int16) {
	v := m.A.Value(registers)
	registers[Acc] *= v
	registers[Acc] = min(registers[Acc], regMax)
	registers[Acc] = max(registers[Acc], regMin)
}

func (m Mul) Accesses() []Reg {
	if reg, ok := m.A.(Reg); ok {
		return []Reg{reg}
	}

	return nil
}

func (m Mul) String() string {
	return fmt.Sprintf("mul %v", m.A)
}
