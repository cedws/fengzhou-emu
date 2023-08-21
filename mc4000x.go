package fengzhouemu

var defaultMC4000XRegisters = map[Reg]Register{
	Null: NullRegister{},
	Acc:  &InternalRegister{},
	X0:   &XbusPinRegister{},
	X1:   &XbusPinRegister{},
	X2:   &XbusPinRegister{},
	X3:   &XbusPinRegister{},
}

type MC4000XProgram [9]Inst

func NewMC4000X(program MC4000XProgram) (*MC, error) {
	registers := make(map[Reg]Register, len(defaultMC4000XRegisters))

	for k, v := range defaultMC4000XRegisters {
		registers[k] = v
	}

	return NewMC(registers, program[:])
}
