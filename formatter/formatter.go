package formatter

import (
	"bytes"
	"monkey/ast"
	"strconv"
)

// FormatAST generates a string representation of the AST, indented for readability.
func FormatAST(node ast.Node) string {
	var buf bytes.Buffer
	formatAstWithDepth(&buf, node, 0)
	return buf.String()
}

// writeIndent writes the specified number of tab characters to the buffer.
func writeIndent(buf *bytes.Buffer, depth int) {
	for i := 0; i < depth; i++ {
		buf.WriteRune('\t') // Use tabs for indentation
	}
}

func formatAstWithDepth(buf *bytes.Buffer, node ast.Node, depth int) {
	switch node := node.(type) {
	case *ast.Program:
		writeIndent(buf, depth)
		buf.WriteString("PROGRAM\n")
		for _, statement := range node.Statements {
			formatAstWithDepth(buf, statement, depth+1)
		}

	case *ast.LetStatement:
		writeIndent(buf, depth)
		buf.WriteString("LET STATEMENT\n")
		writeIndent(buf, depth+1)
		buf.WriteString("(NAME)\n")
		formatAstWithDepth(buf, node.Name, depth+2)
		writeIndent(buf, depth+1)
		buf.WriteString("(VALUE)\n")
		formatAstWithDepth(buf, node.Value, depth+2)

	case *ast.Identifier:
		writeIndent(buf, depth)
		buf.WriteString("IDENTIFIER: ")
		buf.WriteString(node.Value)
		buf.WriteRune('\n')

	case *ast.IntegerLiteral:
		writeIndent(buf, depth)
		buf.WriteString("INTEGER: ")
		buf.WriteString(strconv.Itoa(int(node.Value)))
		buf.WriteRune('\n')

	case *ast.ReturnStatement:
		writeIndent(buf, depth)
		buf.WriteString("RETURN STATEMENT\n")
		if node.ReturnValue != nil {
			writeIndent(buf, depth+1)
			buf.WriteString("(VALUE)\n")
			formatAstWithDepth(buf, node.ReturnValue, depth+2)
		}

	case *ast.ExpressionStatement:
		writeIndent(buf, depth)
		buf.WriteString("EXPRESSION STATEMENT\n")
		formatAstWithDepth(buf, node.Expression, depth+1)

	case *ast.BlockStatement:
		writeIndent(buf, depth)
		buf.WriteString("BLOCK STATEMENT\n")
		for _, statement := range node.Statements {
			formatAstWithDepth(buf, statement, depth+1)
		}

	case *ast.Boolean:
		writeIndent(buf, depth)
		buf.WriteString("BOOLEAN: ")
		buf.WriteString(strconv.FormatBool(node.Value))
		buf.WriteRune('\n')

	case *ast.PrefixExpression:
		writeIndent(buf, depth)
		buf.WriteString("PREFIX EXPRESSION\n")
		writeIndent(buf, depth+1)
		buf.WriteString("OPERATOR: " + node.Operator + "\n")
		writeIndent(buf, depth+1)
		buf.WriteString("RIGHT:\n")
		formatAstWithDepth(buf, node.Right, depth+2)

	case *ast.InfixExpression:
		writeIndent(buf, depth)
		buf.WriteString("INFIX EXPRESSION\n")
		writeIndent(buf, depth+1)
		buf.WriteString("OPERATOR: " + node.Operator + "\n")
		writeIndent(buf, depth+1)
		buf.WriteString("LEFT:\n")
		formatAstWithDepth(buf, node.Left, depth+2)
		writeIndent(buf, depth+1)
		buf.WriteString("RIGHT:\n")
		formatAstWithDepth(buf, node.Right, depth+2)

	case *ast.IfExpression:
		writeIndent(buf, depth)
		buf.WriteString("IF EXPRESSION\n")
		writeIndent(buf, depth+1)
		buf.WriteString("CONDITION:\n")
		formatAstWithDepth(buf, node.Condition, depth+2)
		writeIndent(buf, depth+1)
		buf.WriteString("CONSEQUENCE:\n")
		formatAstWithDepth(buf, node.Consequence, depth+2)
		if node.Alternative != nil {
			writeIndent(buf, depth+1)
			buf.WriteString("ALTERNATIVE:\n")
			formatAstWithDepth(buf, node.Alternative, depth+2)
		}

	case *ast.FunctionLiteral:
		writeIndent(buf, depth)
		buf.WriteString("FUNCTION LITERAL\n")
		writeIndent(buf, depth+1)
		buf.WriteString("PARAMETERS:\n")
		for _, param := range node.Parameters {
			formatAstWithDepth(buf, param, depth+2)
		}
		writeIndent(buf, depth+1)
		buf.WriteString("BODY:\n")
		formatAstWithDepth(buf, node.Body, depth+2)

	case *ast.CallExpression:
		writeIndent(buf, depth)
		buf.WriteString("CALL EXPRESSION\n")
		writeIndent(buf, depth+1)
		buf.WriteString("FUNCTION:\n")
		formatAstWithDepth(buf, node.Function, depth+2)
		writeIndent(buf, depth+1)
		buf.WriteString("ARGUMENTS:\n")
		for _, arg := range node.Arguments {
			formatAstWithDepth(buf, arg, depth+2)
		}

	case *ast.StringLiteral:
		writeIndent(buf, depth)
		buf.WriteString("STRING: ")
		buf.WriteString(node.Value)
		buf.WriteRune('\n')

	case *ast.ArrayLiteral:
		writeIndent(buf, depth)
		buf.WriteString("ARRAY LITERAL\n")
		writeIndent(buf, depth+1)
		buf.WriteString("ELEMENTS:\n")
		for _, element := range node.Elements {
			formatAstWithDepth(buf, element, depth+2)
		}

	case *ast.IndexExpression:
		writeIndent(buf, depth)
		buf.WriteString("INDEX EXPRESSION\n")
		writeIndent(buf, depth+1)
		buf.WriteString("LEFT:\n")
		formatAstWithDepth(buf, node.Left, depth+2)
		writeIndent(buf, depth+1)
		buf.WriteString("INDEX:\n")
		formatAstWithDepth(buf, node.Index, depth+2)

	case *ast.HashLiteral:
		writeIndent(buf, depth)
		buf.WriteString("HASH LITERAL\n")
		writeIndent(buf, depth+1)
		buf.WriteString("PAIRS:\n")
		for key, value := range node.Pairs {
			writeIndent(buf, depth+2)
			buf.WriteString("KEY:\n")
			formatAstWithDepth(buf, key, depth+3)
			writeIndent(buf, depth+2)
			buf.WriteString("VALUE:\n")
			formatAstWithDepth(buf, value, depth+3)
		}
	}
}
