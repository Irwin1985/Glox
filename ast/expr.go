package ast

import (
	"Glox/token"
)

type ExprVisitor interface {
	VisitAssignExpr(expr *Assign) interface{}
	VisitBinaryExpr(expr *Binary) interface{}
	VisitCallExpr(expr *Call) interface{}
	VisitGetExpr(expr *Get) interface{}
	VisitGroupingExpr(expr *Grouping) interface{}
	VisitLiteralExpr(expr *Literal) interface{}
	VisitLogicalExpr(expr *Logical) interface{}
	VisitSetExpr(expr *Set) interface{}
	VisitSuperExpr(expr *Super) interface{}
	VisitThisExpr(expr *This) interface{}
	VisitUnaryExpr(expr *Unary) interface{}
	VisitVariableExpr(expr *Variable) interface{}
}

type Expression interface {
	String() string
	Accept(visitor ExprVisitor) interface{}
}

type Assign struct {
	Name  token.Token
	Value Expression
}

func (expr *Assign) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitAssignExpr(expr)
}
func (expr *Assign) String() string {
	return Beautify(expr)
}

type Variable struct {
	Name token.Token
}

func (expr *Variable) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitVariableExpr(expr)
}
func (expr *Variable) String() string {
	return Beautify(expr)
}

type Unary struct {
	Operator token.Token
	Right    Expression
}

func (expr *Unary) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitUnaryExpr(expr)
}
func (expr *Unary) String() string {
	return Beautify(expr)
}

type Binary struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (expr *Binary) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitBinaryExpr(expr)
}
func (expr *Binary) String() string {
	return Beautify(expr)
}

type Call struct {
	Callee    Expression
	Paren     token.Token
	Arguments []Expression
}

func (expr *Call) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitCallExpr(expr)
}
func (expr *Call) String() string {
	return Beautify(expr)
}

type Get struct {
	Object Expression
	Name   token.Token
}

func (expr *Get) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGetExpr(expr)
}
func (expr *Get) String() string {
	return Beautify(expr)
}

type Set struct {
	Object Expression
	Name   token.Token
	Value  Expression
}

func (expr *Set) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitSetExpr(expr)
}
func (expr *Set) String() string {
	return Beautify(expr)
}

type Grouping struct {
	Expression Expression
}

func (expr *Grouping) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGroupingExpr(expr)
}
func (expr *Grouping) String() string {
	return Beautify(expr)
}

type Literal struct {
	Value interface{}
}

func (expr *Literal) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLiteralExpr(expr)
}
func (expr *Literal) String() string {
	return Beautify(expr)
}

type Logical struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (expr *Logical) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLogicalExpr(expr)
}
func (expr *Logical) String() string {
	return Beautify(expr)
}

type Super struct {
	Keyword token.Token
	Method  token.Token
}

func (expr *Super) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitSuperExpr(expr)
}
func (expr *Super) String() string {
	return Beautify(expr)
}

type This struct {
	Keyword token.Token
}

func (expr *This) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitThisExpr(expr)
}
func (expr *This) String() string {
	return Beautify(expr)
}
