package fengzhouemu

import "fmt"

type (
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
)

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
