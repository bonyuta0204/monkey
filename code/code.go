package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op Opcode) (*Definition, error) {

	definition, ok := definitions[op]

	if !ok {
		return nil, fmt.Errorf("opcode %d is undefined", op)
	}

	return definition, nil

}

func Make(op Opcode, operands ...int) []byte {

	def, ok := definitions[op]

	if !ok {
		return []byte{}
	}

	instructionLength := 1

	for _, width := range def.OperandWidths {
		instructionLength += width
	}

	instruction := make([]byte, instructionLength)

	instruction[0] = byte(op)

	offset := 1

	for i, o := range operands {
		width := def.OperandWidths[i]

		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}

		offset += width
	}

	return instruction

}
