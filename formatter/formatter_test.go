package formatter

import (
	"monkey/lexer"
	"monkey/parser"
	"testing"
)

func TestFormatAST(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input: `
			let x = 5;
			let y = 10;
			if (x > y) {
				return x + y;
			}
			`,
			expected: `PROGRAM
	LET STATEMENT
		(NAME)
			IDENTIFIER: x
		(VALUE)
			INTEGER: 5
	LET STATEMENT
		(NAME)
			IDENTIFIER: y
		(VALUE)
			INTEGER: 10
	EXPRESSION STATEMENT
		IF EXPRESSION
			CONDITION:
				INFIX EXPRESSION
					OPERATOR: >
					LEFT:
						IDENTIFIER: x
					RIGHT:
						IDENTIFIER: y
			CONSEQUENCE:
				BLOCK STATEMENT
					RETURN STATEMENT
						(VALUE)
							INFIX EXPRESSION
								OPERATOR: +
								LEFT:
									IDENTIFIER: x
								RIGHT:
									IDENTIFIER: y
`,
		},
		{
			input: `let add = fn(a, b) { return a + b; }; add(1, 2);`,
			expected: `PROGRAM
	LET STATEMENT
		(NAME)
			IDENTIFIER: add
		(VALUE)
			FUNCTION LITERAL
				PARAMETERS:
					IDENTIFIER: a
					IDENTIFIER: b
				BODY:
					BLOCK STATEMENT
						RETURN STATEMENT
							(VALUE)
								INFIX EXPRESSION
									OPERATOR: +
									LEFT:
										IDENTIFIER: a
									RIGHT:
										IDENTIFIER: b
	EXPRESSION STATEMENT
		CALL EXPRESSION
			FUNCTION:
				IDENTIFIER: add
			ARGUMENTS:
				INTEGER: 1
				INTEGER: 2
`,
		},
		{
			input: `[1, true, "hello"]; { "one": 1, "two": 2 }; !false; -5; a[1];`,
			expected: `PROGRAM
	EXPRESSION STATEMENT
		ARRAY LITERAL
			ELEMENTS:
				INTEGER: 1
				BOOLEAN: true
				STRING: hello
	EXPRESSION STATEMENT
		HASH LITERAL
			PAIRS:
				KEY:
					STRING: one
				VALUE:
					INTEGER: 1
				KEY:
					STRING: two
				VALUE:
					INTEGER: 2
	EXPRESSION STATEMENT
		PREFIX EXPRESSION
			OPERATOR: !
			RIGHT:
				BOOLEAN: false
	EXPRESSION STATEMENT
		PREFIX EXPRESSION
			OPERATOR: -
			RIGHT:
				INTEGER: 5
	EXPRESSION STATEMENT
		INDEX EXPRESSION
			LEFT:
				IDENTIFIER: a
			INDEX:
				INTEGER: 1
`,
		},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()

		actual := FormatAST(program)

		if actual != tt.expected {
			t.Errorf("test[%d]: expected\n%q\n, got \n%q", i, tt.expected, actual)
		}
	}
}
