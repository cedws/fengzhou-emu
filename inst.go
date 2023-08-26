package fengzhouemu

type ConditionType int

const (
	Always ConditionType = iota
	Once
	Enable
	Disable
)

type condition struct {
	Inst
	c ConditionType
}

type labelled struct {
	Inst
	l string
}

func (l labelled) Label() string {
	return l.l
}

func Label(l string, i Inst) Inst {
	return labelled{
		Inst: i,
		l:    l,
	}
}

func (c condition) Condition() ConditionType {
	return c.c
}

func Condition(c ConditionType, i Inst) Inst {
	return condition{
		Inst: i,
		c:    c,
	}
}

// Operand represents a generic operand that can be either a Register or immediate value.
type Operand interface {
	Value(map[Reg]Register) int16
}

// Inst represents a single instruction.
type Inst interface {
	Validate() error
	Execute(map[Reg]Register)
	Cost() int
	Accesses() []Reg
	String() string
	Label() string
	Condition() ConditionType
}
