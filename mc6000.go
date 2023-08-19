package fengzhouemu

var defaultMC6000Registers = map[Reg]int16{
	Acc: 0,
	Dat: 0,
	P0:  0,
	P1:  0,
	X0:  0,
	X1:  0,
	X2:  0,
	X3:  0,
}

type MC6000Program [14]Inst

func NewMC6000(program MC6000Program) (*MC, error) {
	registers := make(map[Reg]int16)

	for k, v := range defaultMC6000Registers {
		registers[k] = v
	}

	return NewMC(registers, program[:])
}
