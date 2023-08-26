package fengzhouemu

func defaultMC4000XRegisters() map[Reg]Register {
	return map[Reg]Register{
		Null:  NullRegister{},
		Acc:   &InternalRegister{},
		X0:    &XbusPinRegister{},
		X1:    &XbusPinRegister{},
		X2:    &XbusPinRegister{},
		X3:    &XbusPinRegister{},
		flags: &InternalRegister{},
		ip:    &InternalRegister{},
	}
}

type MC4000XProgram [9]Inst

func NewMC4000X(program MC4000XProgram) (*MC, error) {
	return NewMC(defaultMC4000XRegisters(), program[:])
}
