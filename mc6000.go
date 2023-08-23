package fengzhouemu

var defaultMC6000Registers = map[Reg]Register{
	Null:  NullRegister{},
	Acc:   &InternalRegister{},
	Dat:   &InternalRegister{},
	P0:    &SimplePinRegister{},
	P1:    &SimplePinRegister{},
	X0:    &XbusPinRegister{},
	X1:    &XbusPinRegister{},
	X2:    &XbusPinRegister{},
	X3:    &XbusPinRegister{},
	flags: &InternalRegister{},
}

type MC6000Program [14]Inst

func NewMC6000(program MC6000Program) (*MC, error) {
	registers := make(map[Reg]Register, len(defaultMC6000Registers))

	for k, v := range defaultMC6000Registers {
		registers[k] = v
	}

	return NewMC(registers, program[:])
}
