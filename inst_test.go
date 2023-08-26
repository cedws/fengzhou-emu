package fengzhouemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstLabelAndCondition(t *testing.T) {
	n := Nop{}
	assert.Equal(t, "", n.Label())

	nl := Label("start", n)
	assert.Equal(t, "start", nl.Label())

	assert.Equal(t, Always, n.Condition())
	nl = Condition(Disable, n)
	assert.Equal(t, Disable, nl.Condition())
}

func TestInstInvalidLabel(t *testing.T) {
	n := Label("hello;", Nop{})
	assert.Equal(t, "hello;", n.Label())

	_, err := NewMC(newRegisters(), []Inst{n})
	assert.ErrorIs(t, err, InvalidLabelNameErr{"hello;"})
}
