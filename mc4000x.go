package fengzhouemu

var defaultMC4000XRegisters = map[Reg]int16{
	Acc: 0,
	X0:  0,
	X1:  0,
	X2:  0,
	X3:  0,
}

type MC4000XProgram [9]Inst

func NewMC4000X(program MC4000XProgram) (*MC, error) {
	registers := make(map[Reg]int16, len(defaultMC4000XRegisters))

	for k, v := range defaultMC4000XRegisters {
		registers[k] = v
	}

	return NewMC(registers, program[:])
}
