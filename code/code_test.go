package code

import (
	"testing"
)

func TestMake(t *testing.T) {
	tests := []struct {
		op       Opcode
		operands []int
		expected []byte
	}{
		{
			OpConstant,
			[]int{65534},
			[]byte{byte(OpConstant), 255, 254},
		},
		{
			OpAdd,
			[]int{},
			[]byte{byte(OpAdd)},
		},
		{
			OpPop,
			[]int{},
			[]byte{byte(OpPop)},
		},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)

		if len(tt.expected) != len(instruction) {

			t.Errorf("instruction has long length. want=%d, got=%d", len(tt.expected), len(instruction))

			for i, b := range tt.expected {
				if b != instruction[i] {
					t.Errorf("wrong byte at %d, want=%d, got %d", i, b, instruction[i])
				}
			}

		}
	}
}

func TestInstructionsString(t *testing.T) {
	instructions := []Instructions{

		Make(OpConstant, 1),
		Make(OpConstant, 2),
		Make(OpConstant, 65535),
		Make(OpAdd),
		Make(OpPop),
	}

	expected := `0000 OpConstant 1
0003 OpConstant 2
0006 OpConstant 65535
0009 OpAdd
0010 OpPop
`

	concatted := Instructions{}

	for _, ins := range instructions {
		concatted = append(concatted, ins...)
	}

	if concatted.String() != expected {
		t.Errorf("instructions wrongly formatted.\nwant=%q\ngot=%q", expected, concatted.String())
	}
}

func TestReadOperands(t *testing.T) {
	tests := []struct {
		op        Opcode
		operands  []int
		bytesRead int
	}{
		{OpConstant, []int{65535}, 2},
		{OpAdd, []int{}, 0},
	}

	for _, tt := range tests {

		instruction := Make(tt.op, tt.operands...)

		def, err := Lookup(byte(tt.op))

		if err != nil {
			t.Fatalf("definition not found %q\n", err)
		}

		operandsRead, n := ReadOperands(def, instruction[1:])

		if n != tt.bytesRead {
			t.Fatalf("n wrong. want=%d, got=%d", tt.bytesRead, n)
		}

		for i, want := range tt.operands {
			if operandsRead[i] != want {
				t.Errorf("operant wrong.want=%d, got=%d", want, operandsRead[i])
			}
		}
	}

}
