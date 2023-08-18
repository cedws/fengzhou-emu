package fengzhouemu

import "fmt"

type OperandType int

const (
	Immediate OperandType = iota
	Register
)

type (
	Imm int16
	Reg int16
)

func (i Imm) Type() OperandType {
	return Immediate
}

func (i Imm) Value(mc *MC4000) int16 {
	return int16(i)
}

const (
	Acc Reg = iota
	P0
	P1
	X0
	X1
	X2
	X3
)

func (r Reg) Type() OperandType {
	return Register
}

func (r Reg) Value(mc *MC4000) int16 {
	return mc.registers[r]
}

type Inst interface {
	Run(*MC4000)
	Cost() int
	Accesses() []Reg
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
		program: program,
		registers: map[Reg]int16{
			Acc: 0,
			P0:  0,
			P1:  0,
			X0:  0,
			X1:  0,
		},
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

	inst.Run(mc)

	mc.power += inst.Cost()
	mc.ip++
}
