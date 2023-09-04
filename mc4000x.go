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

type MC4000X struct {
	*MC
}

type MC4000XProgram [9]Inst

func NewMC4000X() *MC4000X {
	return &MC4000X{
		NewMC(defaultMC4000XRegisters()),
	}
}

func (mc *MC4000X) Load(program MC4000XProgram) error {
	return mc.MC.Load(program[:])
}

func (mc *MC4000X) X0() Register {
	return mc.reg[X0]
}

func (mc *MC4000X) X1() Register {
	return mc.reg[X1]
}

func (mc *MC4000X) X2() Register {
	return mc.reg[X2]
}

func (mc *MC4000X) X3() Register {
	return mc.reg[X3]
}
