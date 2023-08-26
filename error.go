package fengzhouemu

import "fmt"

type (
	ProgramTooLargeErr struct {
		Len int
	}
	NumberTooSmallErr struct {
		Imm Imm
	}
	NumberTooLargeErr struct {
		Imm Imm
	}
	InvalidRegisterErr struct {
		Reg string
	}
	InvalidLabelNameErr struct {
		Label string
	}
	LabelNotDefinedErr struct {
		Label string
	}
	LabelAlreadyDefinedErr struct {
		Label string
	}
)

func (e ProgramTooLargeErr) Error() string {
	return fmt.Sprintf("program too large (%v)", e.Len)
}

func (e NumberTooSmallErr) Error() string {
	return fmt.Sprintf("number too small (%v)", e.Imm)
}

func (e NumberTooLargeErr) Error() string {
	return fmt.Sprintf("number too large (%v)", e.Imm)
}

func (e InvalidRegisterErr) Error() string {
	return fmt.Sprintf("invalid register (%v)", e.Reg)
}

func (e InvalidLabelNameErr) Error() string {
	return fmt.Sprintf("invalid label name (%v)", e.Label)
}

func (e LabelNotDefinedErr) Error() string {
	return fmt.Sprintf("label not defined (%v)", e.Label)
}

func (e LabelAlreadyDefinedErr) Error() string {
	return fmt.Sprintf("label already defined (%v)", e.Label)
}
