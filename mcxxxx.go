package fengzhouemu

import (
	"regexp"
)

const (
	simplePinRegMin = 0
	simplePinRegMax = 100
	regMin          = -999
	regMax          = 999
)

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

func (i Imm) Value(registers map[Reg]Register) int16 {
	return int16(i)
}

// MC is a generic microcontroller.
type MC struct {
	program  []Inst
	executed map[int]bool

	power     int
	ip        int
	registers map[Reg]Register
}

// NewMC creates a new generic microcontroller with the given registers and no limits on program size.
func NewMC(registers map[Reg]Register, program []Inst) (*MC, error) {
	mc := &MC{
		program:   program,
		executed:  make(map[int]bool, len(program)),
		registers: registers,
	}

	if err := mc.Validate(program); err != nil {
		return nil, err
	}

	return mc, nil
}

func (mc *MC) fetch() Inst {
	if len(mc.program) == 0 {
		// variable size program is empty
		return nil
	}

	inst := mc.program[mc.ip]
	if inst != nil {
		mc.ip++

		if mc.ip >= len(mc.program) {
			mc.ip = 0
		}

		return inst
	}

	if mc.ip == 0 {
		// fixed size program is empty
		return nil
	}

	// reached end of fixed size program, or instruction gap, or nil instruction
	mc.ip = 0
	return mc.program[mc.ip]
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
			// programs may not access flags register
			if reg == flags {
				return InvalidRegisterErr{reg.String()}
			}

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
	var inst Inst

	for {
		inst = mc.fetch()
		if inst == nil {
			return
		}

		flags := mc.registers[flags].Read()

		switch inst.Condition() {
		case Once:
			// already executed
			if mc.executed[mc.ip] {
				continue
			}

			mc.executed[mc.ip] = true
		case Enable:
			if flags&testFlag == 0 || flags&enableFlag == 0 {
				continue
			}
		case Disable:
			if flags&testFlag == 0 || flags&enableFlag != 0 {
				continue
			}
		}

		break
	}

	inst.Execute(mc.registers)
	mc.power += inst.Cost()
}
