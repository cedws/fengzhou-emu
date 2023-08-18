package fengzhouemu

type OperandType int

const (
	Immediate OperandType = iota
	Register
)

type (
	Imm int16
	Reg int16
)

func (i Imm) Type() OperandType {
	return Immediate
}

func (i Imm) Value(registers map[Reg]int16) int16 {
	return int16(i)
}

const (
	Acc Reg = iota
	P0
	P1
	X0
	X1
	X2
	X3
)

func (r Reg) Type() OperandType {
	return Register
}

func (r Reg) Value(registers map[Reg]int16) int16 {
	return registers[r]
}
