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

type MC4000 struct {
	*MC
}

type MC4000Program [9]Inst

func NewMC4000() *MC4000 {
	return &MC4000{
		NewMC(defaultMC4000Registers()),
	}
}

func (mc *MC4000) Load(program MC4000Program) error {
	return mc.MC.Load(program[:])
}

func (mc *MC4000) P0() Register {
	return mc.reg[P0]
}

func (mc *MC4000) P1() Register {
	return mc.reg[P1]
}

func (mc *MC4000) X0() Register {
	return mc.reg[X0]
}

func (mc *MC4000) X1() Register {
	return mc.reg[X1]
}
