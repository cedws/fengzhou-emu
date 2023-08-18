package fengzhouemu

const (
	regMin = -999
	regMax = 999
)

type Operand interface {
	Type() OperandType
	Value(map[Reg]int16) int16
}

type Inst interface {
	Execute(map[Reg]int16)
	Cost() int
	Accesses() []Reg
}

type Nop struct{}

func (n Nop) Cost() int {
	return 1
}

func (n Nop) Execute(registers map[Reg]int16) {}

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

func (m Mov) Execute(registers map[Reg]int16) {
	v := m.A.Value(registers)
	registers[m.B] = v
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

func (m Add) Execute(registers map[Reg]int16) {
	v := m.A.Value(registers)
	registers[Acc] += v
	registers[Acc] = min(registers[Acc], regMax)
	registers[Acc] = max(registers[Acc], regMin)
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

func (m Sub) Execute(registers map[Reg]int16) {
	v := m.A.Value(registers)
	registers[Acc] -= v
	registers[Acc] = min(registers[Acc], regMax)
	registers[Acc] = max(registers[Acc], regMin)
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

func (m Mul) Execute(registers map[Reg]int16) {
	v := m.A.Value(registers)
	registers[Acc] *= v
	registers[Acc] = min(registers[Acc], regMax)
	registers[Acc] = max(registers[Acc], regMin)
}

func (m Mul) Accesses() []Reg {
	if m.A.Type() == Register {
		return []Reg{m.A.(Reg)}
	}
	return nil
}
