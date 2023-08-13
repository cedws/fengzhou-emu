package fengzhouemu

const (
	xbusMin = -999
	xbusMax = 999
)

type Nop struct{}

func (n Nop) Cost() int {
	return 1
}

func (n Nop) Run(mc *MC4000) {}

type Mov struct {
	A interface {
		Value(*MC4000) int16
	}
	B Reg
}

func (n Mov) Cost() int {
	return 1
}

func (m Mov) Run(mc *MC4000) {
	v := m.A.Value(mc)
	mc.registers[m.B] = v
}

type Add struct {
	A interface {
		Value(*MC4000) int16
	}
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

type Sub struct {
	A interface {
		Value(*MC4000) int16
	}
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

type Mul struct {
	A interface {
		Value(*MC4000) int16
	}
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
