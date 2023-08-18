package fengzhouemu

import "fmt"

var defaultMC4000Registers = map[Reg]int16{
	Acc: 0,
	P0:  0,
	P1:  0,
	X0:  0,
	X1:  0,
}

type MC4000Program [9]Inst

type MC4000 struct {
	program   MC4000Program
	power     int
	ip        byte
	registers map[Reg]int16
}

func NewMC4000(program MC4000Program) (*MC4000, error) {
	mc := &MC4000{
		program:   program,
		registers: make(map[Reg]int16),
	}

	for k, v := range defaultMC4000Registers {
		mc.registers[k] = v
	}

	if err := mc.Validate(program); err != nil {
		return nil, err
	}

	return mc, nil
}

func (mc *MC4000) Validate(program MC4000Program) error {
	for _, inst := range program {
		if inst == nil {
			continue
		}

		accesses := inst.Accesses()

		for _, reg := range accesses {
			if _, ok := mc.registers[reg]; !ok {
				return fmt.Errorf("invalid register %d", reg)
			}
		}
	}

	return nil
}

func (mc *MC4000) Power() int {
	return mc.power
}

func (mc *MC4000) Step() {
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
