package ast

import (
	"Glox/token"
)

type StmtVisitor interface {
	VisitBlockStmt(stmt *BlockStmt) interface{}
	VisitClassStmt(stmt *ClassStmt) interface{}
	VisitExpressionStmt(stmt *ExpressionStmt) interface{}
	VisitFunctionStmt(stmt *FunStmt) interface{}
	VisitIfStmt(stmt *IfStmt) interface{}
	VisitPrintStmt(stmt *PrintStmt) interface{}
	VisitReturnStmt(stmt *ReturnStmt) interface{}
	VisitVarStmt(stmt *VarStmt) interface{}
	VisitWhileStmt(stmt *WhileStmt) interface{}
}

type Statement interface {
	String() string
	Accept(visitor StmtVisitor) interface{}
}

type BlockStmt struct {
	Statements []Statement
}

func (stmt *BlockStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitBlockStmt(stmt)
}
func (stmt *BlockStmt) String() string {
	return Beautify(stmt)
}

type ClassStmt struct {
	Name       token.Token
	Superclass *Variable
	Methods    []Statement
}

func (stmt *ClassStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitClassStmt(stmt)
}
func (stmt *ClassStmt) String() string {
	return Beautify(stmt)
}

type ExpressionStmt struct {
	Expression Expression
}

func (stmt *ExpressionStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitExpressionStmt(stmt)
}
func (stmt *ExpressionStmt) String() string {
	return Beautify(stmt)
}

type FunStmt struct {
	Name   token.Token
	Params []token.Token
	Body   []Statement
}

func (stmt *FunStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitFunctionStmt(stmt)
}
func (stmt *FunStmt) String() string {
	return Beautify(stmt)
}

type IfStmt struct {
	Condition  Expression
	ThenBranch Statement
	ElseBranch Statement
}

func (stmt *IfStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitIfStmt(stmt)
}
func (stmt *IfStmt) String() string {
	return Beautify(stmt)
}

type PrintStmt struct {
	Expression Expression
}

func (stmt *PrintStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitPrintStmt(stmt)
}
func (stmt *PrintStmt) String() string {
	return Beautify(stmt)
}

type ReturnStmt struct {
	Keyword token.Token
	Value   Expression
}

func (stmt *ReturnStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitReturnStmt(stmt)
}
func (stmt *ReturnStmt) String() string {
	return Beautify(stmt)
}

type VarStmt struct {
	Name        token.Token
	Initializer Expression
}

func (stmt *VarStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitVarStmt(stmt)
}
func (stmt *VarStmt) String() string {
	return Beautify(stmt)
}

type WhileStmt struct {
	Condition Expression
	Body      Statement
}

func (stmt *WhileStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitWhileStmt(stmt)
}
func (stmt *WhileStmt) String() string {
	return Beautify(stmt)
}
