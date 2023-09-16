package fengzhouemu

type PinMode int

const (
	Input PinMode = iota
	Output
)

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
	Mode() PinMode
	Wired() bool
	WireTo(*Link)
	Write(int16)
	Read() int16
	Value() int16
}

type NullRegister struct{}

func (n NullRegister) Mode() PinMode {
	return Output
}

func (n NullRegister) Wired() bool {
	return true
}

func (n NullRegister) WireTo(l *Link) {}

func (n NullRegister) Write(int16) {}

func (n NullRegister) Read() int16 {
	return 0
}

func (n NullRegister) Value() int16 {
	return 0
}

type InternalRegister struct {
	value int16
}

func (i *InternalRegister) Mode() PinMode {
	return Output
}

func (i *InternalRegister) Wired() bool {
	return true
}

func (i *InternalRegister) WireTo(l *Link) {}

func (i *InternalRegister) Write(v int16) {
	i.value = v
	i.value = min(i.value, regMax)
	i.value = max(i.value, regMin)
}

func (i *InternalRegister) Read() int16 {
	return i.value
}

func (i *InternalRegister) Value() int16 {
	return i.value
}

type SimplePinRegister struct {
	mode  PinMode
	link  *Link
	value int16
}

func (s *SimplePinRegister) Mode() PinMode {
	return s.mode
}

func (s *SimplePinRegister) Wired() bool {
	return s.link != nil
}

func (s *SimplePinRegister) WireTo(l *Link) {
	if !s.Wired() {
		s.link = l
		s.link.wireTo(s)
	}
}

func (s *SimplePinRegister) Write(v int16) {
	s.mode = Output
	s.value = v
}

func (s *SimplePinRegister) Read() int16 {
	s.mode = Input
	s.value = 0

	if s.link != nil {
		return s.link.read()
	}

	return 0
}

func (s *SimplePinRegister) Value() int16 {
	return s.value
}

type XbusPinRegister struct {
	mode  PinMode
	value int16
}

func (x *XbusPinRegister) Mode() PinMode {
	return x.mode
}

func (x *XbusPinRegister) Write(v int16) {

}

func (x *XbusPinRegister) Wired() bool {
	return true
}

func (x *XbusPinRegister) WireTo(l *Link) {

}

func (x *XbusPinRegister) Read() int16 {
	return 0
}

func (x *XbusPinRegister) Value() int16 {
	return x.value
}

type Link struct {
	pins []Register
}

func (l *Link) wireTo(pin Register) {
	l.pins = append(l.pins, pin)
}

func (l *Link) read() int16 {
	var v int16

	for _, pin := range l.pins {
		if pin.Mode() == Output {
			v = max(v, pin.Value())
		}
	}

	return v
}
