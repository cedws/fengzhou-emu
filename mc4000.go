package fengzhouemu

var defaultMC4000Registers = map[Reg]int16{
	Acc: 0,
	P0:  0,
	P1:  0,
	X0:  0,
	X1:  0,
}

type MC4000Program [9]Inst

func NewMC4000(program MC4000Program) (*MC, error) {
	registers := make(map[Reg]int16, len(defaultMC4000Registers))

	for k, v := range defaultMC4000Registers {
		registers[k] = v
	}

	return NewMC(registers, program[:])
}
