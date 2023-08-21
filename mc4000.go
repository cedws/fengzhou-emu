package fengzhouemu

var defaultMC4000Registers = map[Reg]Register{
	Null: NullRegister{},
	Acc:  &InternalRegister{},
	P0:   &SimplePinRegister{},
	P1:   &SimplePinRegister{},
	X0:   &XbusPinRegister{},
	X1:   &XbusPinRegister{},
}

type MC4000Program [9]Inst

func NewMC4000(program MC4000Program) (*MC, error) {
	registers := make(map[Reg]Register, len(defaultMC4000Registers))

	for k, v := range defaultMC4000Registers {
		registers[k] = v
	}

	return NewMC(registers, program[:])
}
