package fengzhouemu

const (
	xbusMin = -999
	xbusMax = 999
)

type Operand interface {
	Type() OperandType
	Value(*MC4000) int16
}

type Nop struct{}

func (n Nop) Cost() int {
	return 1
}

func (n Nop) Run(mc *MC4000) {}

func (n Nop) Accesses() []Reg {
	return nil
}

type Mov struct {
	A Operand
	B Reg
}

func (n Mov) Cost() int {
	return 1
}

func (m Mov) Run(mc *MC4000) {
	v := m.A.Value(mc)
	mc.registers[m.B] = v
}

func (m Mov) Accesses() []Reg {
	if m.A.Type() == Register {
		return []Reg{m.A.(Reg), m.B}
	}
	return []Reg{m.B}
}

type Add struct {
	A Operand
}

func (n Add) Cost() int {
	return 1
}

func (m Add) Run(mc *MC4000) {
	v := m.A.Value(mc)
	mc.registers[Acc] += v
	mc.registers[Acc] = min(mc.registers[Acc], xbusMax)
	mc.registers[Acc] = max(mc.registers[Acc], xbusMin)
}

func (m Add) Accesses() []Reg {
	if m.A.Type() == Register {
		return []Reg{m.A.(Reg)}
	}
	return nil
}

type Sub struct {
	A Operand
}

func (n Sub) Cost() int {
	return 1
}

func (m Sub) Run(mc *MC4000) {
	v := m.A.Value(mc)
	mc.registers[Acc] -= v
	mc.registers[Acc] = min(mc.registers[Acc], xbusMax)
	mc.registers[Acc] = max(mc.registers[Acc], xbusMin)
}

func (m Sub) Accesses() []Reg {
	if m.A.Type() == Register {
		return []Reg{m.A.(Reg)}
	}
	return nil
}

type Mul struct {
	A Operand
}

func (n Mul) Cost() int {
	return 1
}

func (m Mul) Run(mc *MC4000) {
	v := m.A.Value(mc)
	mc.registers[Acc] *= v
	mc.registers[Acc] = min(mc.registers[Acc], xbusMax)
	mc.registers[Acc] = max(mc.registers[Acc], xbusMin)
}

func (m Mul) Accesses() []Reg {
	if m.A.Type() == Register {
		return []Reg{m.A.(Reg)}
	}
	return nil
}
