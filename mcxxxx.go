package fengzhouemu

import (
	"math"
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
	executed map[int16]bool
	labels   map[string]int16

	power int
	reg   map[Reg]Register
}

// NewMC creates a new generic microcontroller with the given registers and no limits on program size.
func NewMC(reg map[Reg]Register, program []Inst) (*MC, error) {
	mc := &MC{
		program:  program,
		executed: make(map[int16]bool, len(program)),
		labels:   make(map[string]int16, len(program)),
		reg:      reg,
	}

	if err := mc.Validate(program); err != nil {
		return nil, err
	}

	return mc, nil
}

func (mc *MC) jump(label string) {
	// guaranteed due to program validation
	ptr := mc.labels[label]
	mc.reg[ip].Write(ptr)
}

func (mc *MC) fetch() Inst {
	if len(mc.program) == 0 {
		// variable size program is empty
		return nil
	}

	currentIp := mc.reg[ip].Read()
	inst := mc.program[currentIp]

	if inst != nil {
		currentIp++

		if int(currentIp) >= len(mc.program) {
			currentIp = 0
		}

		mc.reg[ip].Write(currentIp)
		return inst
	}

	if currentIp == 0 {
		// fixed size program is empty
		return nil
	}

	// reached end of fixed size program, or instruction gap, or nil instruction
	mc.reg[ip].Write(0)
	return mc.program[0]
}

// Validate checks that the given program is valid for this microcontroller.
func (mc *MC) Validate(program []Inst) error {
	if len(program) >= math.MaxInt16 {
		return ProgramTooLargeErr{len(program)}
	}

	for i, inst := range program {
		if inst == nil {
			continue
		}

		label := inst.Label()
		if label != "" {
			if !validLabel.MatchString(label) {
				return InvalidLabelNameErr{label}
			}

			if _, ok := mc.labels[label]; ok {
				return LabelAlreadyDefinedErr{label}
			}

			mc.labels[label] = int16(i)
		}

		accesses := inst.Accesses()
		for _, reg := range accesses {
			// programs may not access private registers
			switch reg {
			case ip, flags:
				return InvalidRegisterErr{reg.String()}
			}

			if _, ok := mc.reg[reg]; !ok {
				return InvalidRegisterErr{reg.String()}
			}
		}

		if err := inst.Validate(); err != nil {
			return err
		}
	}

	// second pass, validate jmps
	for _, inst := range program {
		if jmp, ok := inst.(Jmp); ok {
			if _, ok := mc.labels[jmp.A]; !ok {
				return LabelNotDefinedErr{jmp.A}
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

		flags := mc.reg[flags].Read()

		switch inst.Condition() {
		case Once:
			ip := mc.reg[ip].Read()

			// already executed
			if mc.executed[ip] {
				continue
			}

			mc.executed[ip] = true
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

	inst.Execute(mc)
	mc.power += inst.Cost()
}
