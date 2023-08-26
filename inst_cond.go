package fengzhouemu

import "fmt"

type Teq struct {
	A Operand
	B Operand
}

func (t Teq) Validate() error {
	if imm, ok := t.A.(Imm); ok {
		if err := imm.Validate(); err != nil {
			return err
		}
	}
	if imm, ok := t.B.(Imm); ok {
		if err := imm.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (t Teq) Cost() int {
	return 1
}

func (t Teq) Execute(mc *MC) {
	currentFlags := mc.reg[flags].Read()
	currentFlags |= testFlag

	if t.A.Value(mc.reg) == t.B.Value(mc.reg) {
		currentFlags |= enableFlag
	} else {
		currentFlags &^= enableFlag
	}

	mc.reg[flags].Write(currentFlags)
}

func (t Teq) Accesses() (registers []Reg) {
	if reg, ok := t.A.(Reg); ok {
		registers = append(registers, reg)
	}
	if reg, ok := t.B.(Reg); ok && reg != t.A {
		registers = append(registers, reg)
	}
	return
}

func (t Teq) String() string {
	return fmt.Sprintf("teq %v %v", t.A, t.B)
}

func (t Teq) Label() string {
	return ""
}

func (t Teq) Condition() ConditionType {
	return Always
}

type Tgt struct {
	A Operand
	B Operand
}

func (t Tgt) Validate() error {
	if imm, ok := t.A.(Imm); ok {
		if err := imm.Validate(); err != nil {
			return err
		}
	}
	if imm, ok := t.B.(Imm); ok {
		if err := imm.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (t Tgt) Cost() int {
	return 1
}

func (t Tgt) Execute(mc *MC) {
	currentFlags := mc.reg[flags].Read()
	currentFlags |= testFlag

	if t.A.Value(mc.reg) > t.B.Value(mc.reg) {
		currentFlags |= enableFlag
	} else {
		currentFlags &^= enableFlag
	}

	mc.reg[flags].Write(currentFlags)
}

func (t Tgt) Accesses() (registers []Reg) {
	if reg, ok := t.A.(Reg); ok {
		registers = append(registers, reg)
	}
	if reg, ok := t.B.(Reg); ok && reg != t.A {
		registers = append(registers, reg)
	}
	return
}

func (t Tgt) String() string {
	return fmt.Sprintf("tgt %v %v", t.A, t.B)
}

func (t Tgt) Label() string {
	return ""
}

func (t Tgt) Condition() ConditionType {
	return Always
}

type Tlt struct {
	A Operand
	B Operand
}

func (t Tlt) Validate() error {
	if imm, ok := t.A.(Imm); ok {
		if err := imm.Validate(); err != nil {
			return err
		}
	}
	if imm, ok := t.B.(Imm); ok {
		if err := imm.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (t Tlt) Cost() int {
	return 1
}

func (t Tlt) Execute(mc *MC) {
	currentFlags := mc.reg[flags].Read()
	currentFlags |= testFlag

	if t.A.Value(mc.reg) < t.B.Value(mc.reg) {
		currentFlags |= enableFlag
	} else {
		currentFlags &^= enableFlag
	}

	mc.reg[flags].Write(currentFlags)
}

func (t Tlt) Accesses() (registers []Reg) {
	if reg, ok := t.A.(Reg); ok {
		registers = append(registers, reg)
	}
	if reg, ok := t.B.(Reg); ok && reg != t.A {
		registers = append(registers, reg)
	}
	return
}

func (t Tlt) String() string {
	return fmt.Sprintf("tlt %v %v", t.A, t.B)
}

func (t Tlt) Label() string {
	return ""
}

func (t Tlt) Condition() ConditionType {
	return Always
}

type Tcp struct {
	A Operand
	B Operand
}

func (t Tcp) Validate() error {
	if imm, ok := t.A.(Imm); ok {
		if err := imm.Validate(); err != nil {
			return err
		}
	}
	if imm, ok := t.B.(Imm); ok {
		if err := imm.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (t Tcp) Cost() int {
	return 1
}

func (t Tcp) Execute(mc *MC) {
	currentFlags := mc.reg[flags].Read()

	switch {
	case t.A.Value(mc.reg) > t.B.Value(mc.reg):
		currentFlags |= testFlag
		currentFlags |= enableFlag
	case t.A.Value(mc.reg) < t.B.Value(mc.reg):
		currentFlags |= testFlag
		currentFlags &^= enableFlag
	default:
		currentFlags &^= testFlag
	}

	mc.reg[flags].Write(currentFlags)
}

func (t Tcp) Accesses() (registers []Reg) {
	if reg, ok := t.A.(Reg); ok {
		registers = append(registers, reg)
	}
	if reg, ok := t.B.(Reg); ok && reg != t.A {
		registers = append(registers, reg)
	}
	return
}

func (t Tcp) String() string {
	return fmt.Sprintf("tcp %v %v", t.A, t.B)
}

func (t Tcp) Label() string {
	return ""
}

func (t Tcp) Condition() ConditionType {
	return Always
}
