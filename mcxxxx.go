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
	program []Inst
	labels  map[string]int16

	power    int
	executed map[int16]bool
	reg      map[Reg]Register
}

// NewMC creates a new generic microcontroller with the given registers and no limits on program size.
func NewMC(reg map[Reg]Register) *MC {
	return &MC{
		reg: reg,
	}
}

// Load loads the given program into the microcontroller.
func (mc *MC) Load(program []Inst) error {
	mc.executed = make(map[int16]bool, len(program))
	mc.labels = make(map[string]int16, len(program))

	if err := mc.Validate(program); err != nil {
		return err
	}

	mc.program = program

	return nil
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
				return InvalidRegisterErr{reg}
			}

			register, ok := mc.reg[reg]
			if !ok {
				return InvalidRegisterErr{reg}
			}

			if !register.Wired() {
				return PinNotConnectedErr{reg}
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
	var (
		inst Inst
		ip   int16
	)

	for {
		inst, ip = mc.fetch()
		if inst == nil {
			return
		}

		switch inst.Condition() {
		case Once:
			// already executed
			if mc.executed[ip] {
				continue
			}

			mc.executed[ip] = true
		case Enable:
			flags := mc.reg[flags].Read()

			if flags&testFlag == 0 || flags&enableFlag == 0 {
				continue
			}
		case Disable:
			flags := mc.reg[flags].Read()

			if flags&testFlag == 0 || flags&enableFlag != 0 {
				continue
			}
		}

		break
	}

	inst.Execute(mc)
	mc.power += inst.Cost()
}

func (mc *MC) jump(label string) {
	// guaranteed due to program validation
	ptr := mc.labels[label]
	mc.reg[ip].Write(ptr)
}

func (mc *MC) fetch() (Inst, int16) {
	if len(mc.program) == 0 {
		// variable size program is empty
		return nil, 0
	}

	currentIP := mc.reg[ip].Read()
	inst := mc.program[currentIP]

	if inst != nil {
		currentIP++

		if int(currentIP) >= len(mc.program) {
			currentIP = 0
		}

		mc.reg[ip].Write(currentIP)
		return inst, currentIP
	}

	if currentIP == 0 {
		// fixed size program is empty
		return nil, currentIP
	}

	// reached end of fixed size program, or instruction gap, or nil instruction
	mc.reg[ip].Write(0)
	return mc.program[0], 0
}
