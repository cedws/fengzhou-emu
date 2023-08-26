package fengzhouemu

const (
	testFlag = 1 << iota
	enableFlag
)

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
	// private registers, not accessible
	ip
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
	case ip, flags:
		fallthrough
	default:
		return "(unknown register)"
	}
}

func (r Reg) Value(registers map[Reg]Register) int16 {
	return registers[r].Read()
}

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
