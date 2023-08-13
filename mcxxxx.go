package fengzhouemu

type (
	Imm int16
	Reg int16
)

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

func (r Reg) Value(mc *MC4000) int16 {
	return mc.registers[r]
}

type Inst interface {
	Run(*MC4000)
	Cost() int
}

type MC4000Program [9]Inst

type MC4000 struct {
	program   MC4000Program
	power     int
	ip        byte
	registers map[Reg]int16
}

func NewMC4000(program MC4000Program) (*MC4000, error) {
	return &MC4000{
		program: program,
		registers: map[Reg]int16{
			Acc: 0,
			P0:  0,
			P1:  0,
			X0:  0,
			X1:  0,
		},
	}, nil
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
