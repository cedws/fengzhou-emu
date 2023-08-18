package fengzhouemu

var defaultMC4000XRegisters = map[Reg]int16{
	Acc: 0,
	X0:  0,
	X1:  0,
	X2:  0,
	X3:  0,
}

type MC4000XProgram [9]Inst

type MC4000X struct {
	program   MC4000XProgram
	power     int
	ip        byte
	registers map[Reg]int16
}

func NewMC4000X(program MC4000XProgram) (*MC4000X, error) {
	mc := &MC4000X{
		program:   program,
		registers: make(map[Reg]int16),
	}

	for k, v := range defaultMC4000XRegisters {
		mc.registers[k] = v
	}

	if err := mc.Validate(program); err != nil {
		return nil, err
	}

	return mc, nil
}

func (mc *MC4000X) Validate(program MC4000XProgram) error {
	for _, inst := range program {
		if inst == nil {
			continue
		}

		accesses := inst.Accesses()

		for _, reg := range accesses {
			if _, ok := mc.registers[reg]; !ok {
				return InvalidRegisterErr{reg.String()}
			}
		}
	}

	return nil
}

func (mc *MC4000X) Power() int {
	return mc.power
}

func (mc *MC4000X) Step() {
	inst := mc.program[mc.ip]
	if inst == nil {
		if mc.ip == 0 {
			return
		}

		mc.ip = 0
		inst = mc.program[mc.ip]
	}

	inst.Execute(mc.registers)

	mc.power += inst.Cost()
	mc.ip++
}
