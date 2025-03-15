package compiler

import (
	"fmt"
	"monkey/ast"
	"monkey/code"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"testing"
)

type compilerTestcase struct {
	input                string
	expectedConstatns    []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []compilerTestcase{
		{
			input:             "1 + 2",
			expectedConstatns: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpAdd),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "1; 2",
			expectedConstatns: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpPop),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpPop),
			},
		},
	}

	runCompileTests(t, tests)
}

func runCompileTests(t *testing.T, tests []compilerTestcase) {
	t.Helper()

	for _, tt := range tests {

		program := parse(tt.input)

		compiler := New()

		err := compiler.Compile(program)

		if err != nil {
			t.Fatalf("compile error: %s", err)
		}

		bytecode := compiler.Bytecode()

		err = testInstructions(tt.expectedInstructions, bytecode.Instructions)

		if err != nil {
			t.Fatalf("testInstructions failed. %s", err)
		}

		err = testConstants(tt.expectedConstatns, bytecode.Constants)

		if err != nil {
			t.Fatalf("testConstants failed. %s", err)
		}

	}

}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)

	return p.ParseProgram()
}

func testInstructions(expected []code.Instructions, actual code.Instructions) error {

	concatted := concatInstrucation(expected)

	if len(concatted) != len(actual) {

		return fmt.Errorf("wrong instrction length.\nwant=%q\ngot=%q", concatted, actual)
	}

	for i, ins := range concatted {
		if ins != actual[i] {
			return fmt.Errorf("wrong instrction at %d.\nwant=%q\ngot=%q", i, ins, actual[i])
		}
	}

	return nil

}

func concatInstrucation(s []code.Instructions) code.Instructions {
	out := code.Instructions{}

	for _, ins := range s {
		out = append(out, ins...)
	}

	return out
}

func testConstants(expected []interface{}, actual []object.Object) error {

	if len(expected) != len(actual) {
		return fmt.Errorf("wrong number of constants. got=%d, want=%d", len(actual), len(expected))
	}

	for i, constant := range expected {
		switch constant := constant.(type) {
		case int:
			err := testIntegerObject(int64(constant), actual[i])

			if err != nil {
				return fmt.Errorf("constant %d - testIntegerObject failed: %s", i, err)
			}
		}
	}

	return nil
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)

	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("Object has wrong value. got=%d, want=%d", result.Value, expected)
	}

	return nil
}
