package interpreter

import (
	"Glox/token"
	"fmt"
)

type Environment struct {
	enclosing *Environment
	values    map[string]interface{}
}

func NewEnvironment() *Environment {
	e := &Environment{
		values: make(map[string]interface{}),
	}
	return e
}

func NewEnclosedEnvironment(enclosing *Environment) *Environment {
	e := NewEnvironment()
	e.enclosing = enclosing
	return e
}

func (e *Environment) Define(name string, value interface{}) {
	e.values[name] = value
}

func (e *Environment) Get(name token.Token) interface{} {
	if value, ok := e.values[name.Lexeme]; ok {
		return value
	}
	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}
	panic(fmt.Errorf("%s Undefined variable '%s'", name.ToString(), name.Lexeme))
}

func (e *Environment) Assign(name token.Token, value interface{}) {
	if _, ok := e.values[name.Lexeme]; ok {
		e.values[name.Lexeme] = value
	}
	if e.enclosing != nil {
		e.enclosing.Assign(name, value)
	}
	panic(fmt.Errorf("%s Undefined variable '%s'", name.ToString(), name.Lexeme))
}
