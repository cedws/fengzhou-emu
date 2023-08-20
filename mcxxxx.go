package fengzhouemu

import "regexp"

var validLabel = regexp.MustCompile(`^\w*$`)

type (
	Imm int16
	Reg int16
)

func (i Imm) Validate() error {
	if i > regMax {
		return NumberTooLargeErr{i}
	}

	if i < regMin {
		return NumberTooSmallErr{i}
	}

	return nil
}

func (i Imm) Value(registers map[Reg]int16) int16 {
	return int16(i)
}

const (
	Null Reg = iota
	Acc
	Dat
	P0
	P1
	X0
	X1
	X2
	X3
)

func (r Reg) String() string {
	switch r {
	case Null:
		return "null"
	case Acc:
		return "acc"
	case Dat:
		return "dat"
	case P0:
		return "p0"
	case P1:
		return "p1"
	case X0:
		return "x0"
	case X1:
		return "x1"
	case X2:
		return "x2"
	case X3:
		return "x3"
	default:
		return "(unknown register)"
	}
}

func (r Reg) Value(registers map[Reg]int16) int16 {
	return registers[r]
}

// MC is a generic microcontroller.
type MC struct {
	program   []Inst
	registers map[Reg]int16

	power int
	ip    byte
}

// NewMC creates a new generic microcontroller with the given registers and no limits on program size.
func NewMC(registers map[Reg]int16, program []Inst) (*MC, error) {
	mc := &MC{
		registers: registers,
		program:   program,
	}

	if err := mc.Validate(program); err != nil {
		return nil, err
	}

	return mc, nil
}

// Validate checks that the given program is valid for this microcontroller.
func (mc *MC) Validate(program []Inst) error {
	for _, inst := range program {
		if inst == nil {
			continue
		}

		label := inst.Label()
		if !validLabel.MatchString(label) {
			return InvalidLabelNameErr{label}
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

// Power returns the power consumed by this microcontroller so far.
func (mc *MC) Power() int {
	return mc.power
}

// Step executes the next instruction in the program.
func (mc *MC) Step() {
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
