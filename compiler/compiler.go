package compiler

import (
	"monkey/ast"
	"monkey/code"
	"monkey/object"
)

type Compiler struct {
	instrctions code.Instructions
	constants   []object.Object
}

func New() *Compiler {

	return &Compiler{
		instrctions: code.Instructions{},
		constants:   []object.Object{},
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		instrctions: c.instrctions,
		constants:   c.constants,
	}
}

type Bytecode struct {
	instrctions code.Instructions
	constants   []object.Object
}
