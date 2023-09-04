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
		ip:    &InternalRegister{},
		flags: &InternalRegister{},
	}
}

type MC6000 struct {
	*MC
}

type MC6000Program [14]Inst

func NewMC6000() *MC6000 {
	return &MC6000{
		NewMC(defaultMC6000Registers()),
	}
}

func (mc *MC6000) Load(program MC6000Program) error {
	return mc.MC.Load(program[:])
}

func (mc *MC6000) P0() Register {
	return mc.reg[P0]
}

func (mc *MC6000) P1() Register {
	return mc.reg[P1]
}

func (mc *MC6000) X0() Register {
	return mc.reg[X0]
}

func (mc *MC6000) X1() Register {
	return mc.reg[X1]
}

func (mc *MC6000) X2() Register {
	return mc.reg[X2]
}

func (mc *MC6000) X3() Register {
	return mc.reg[X3]
}
