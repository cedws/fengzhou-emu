package fengzhouemu

type (
	Imm int16
	Reg int16
)

func (i Imm) Validate() error {
	if i > regMax {
		return NumberTooLargeErr{}
	}

	if i < regMin {
		return NumberTooSmallErr{}
	}

	return nil
}

func (i Imm) Value(registers map[Reg]int16) int16 {
	return int16(i)
}

const (
	Acc Reg = iota
	Dat
	P0
	P1
	X0
	X1
	X2
	X3
)

func (r Reg) String() string {
	switch r {
	case Acc:
		return "acc"
	case Dat:
		return "dat"
	case P0:
		return "p0"
	case P1:
		return "p1"
	case X0:
		return "x0"
	case X1:
		return "x1"
	case X2:
		return "x2"
	case X3:
		return "x3"
	default:
		return ""
	}
}

func (r Reg) Value(registers map[Reg]int16) int16 {
	return registers[r]
}
