package fengzhouemu

func defaultMC4000Registers() map[Reg]Register {
	return map[Reg]Register{
		Null:  NullRegister{},
		Acc:   &InternalRegister{},
		P0:    &SimplePinRegister{},
		P1:    &SimplePinRegister{},
		X0:    &XbusPinRegister{},
		X1:    &XbusPinRegister{},
		flags: &InternalRegister{},
		ip:    &InternalRegister{},
	}
}

type MC4000Program [9]Inst

func NewMC4000(program MC4000Program) (*MC, error) {
	return NewMC(defaultMC4000Registers(), program[:])
}
