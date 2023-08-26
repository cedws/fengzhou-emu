package fengzhouemu

func defaultMC6000Registers() map[Reg]Register {
	return map[Reg]Register{
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
}

type MC6000Program [14]Inst

func NewMC6000(program MC6000Program) (*MC, error) {
	return NewMC(defaultMC6000Registers(), program[:])
}
