package fengzhouemu

import (
	"regexp"
)

type Register interface {
	Write(int16)
	Read() int16
}

type NullRegister struct{}

func (n NullRegister) Write(int16) {}

func (n NullRegister) Read() int16 {
	return 0
}

type InternalRegister struct {
	value int16
}

func (i *InternalRegister) Write(v int16) {
	i.value = v
	i.value = min(i.value, regMax)
	i.value = max(i.value, regMin)
}

func (i *InternalRegister) Read() int16 {
	return i.value
}

type SimplePinRegister struct {
	value int16
}

func (s *SimplePinRegister) Write(v int16) {
	s.value = v
	s.value = min(s.value, simplePinRegMax)
	s.value = max(s.value, simplePinRegMin)
}

func (s *SimplePinRegister) Read() int16 {
	return s.value
}

type XbusPinRegister struct {
	valueCh chan int16
}

func (x *XbusPinRegister) Write(v int16) {
	x.valueCh <- v
}

func (x *XbusPinRegister) Read() int16 {
	return <-x.valueCh
}

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

const (
	// internal registers
	Null Reg = iota
	Acc
	Dat
	// simple pin registers
	P0
	P1
	// xbus pin registers
	X0
	X1
	X2
	X3
	// flags register, not accessible
	flags
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
	case flags:
		fallthrough
	default:
		return "(unknown register)"
	}
}

func (r Reg) Value(registers map[Reg]Register) int16 {
	return registers[r].Read()
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

		if inst.Condition() == Once {
			// already executed
			if mc.executed[mc.ip] {
				continue
			}

			mc.executed[mc.ip] = true
		}

		break
	}

	inst.Execute(mc.registers)
	mc.power += inst.Cost()
}
