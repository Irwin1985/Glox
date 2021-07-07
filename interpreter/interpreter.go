package interpreter

import (
	"Glox/ast"
	"Glox/token"
	"fmt"
)

type Interpreter struct {
	environment *Environment
}

func NewInterpreter() *Interpreter {
	i := &Interpreter{}
	i.environment = NewEnvironment()
	return i
}

func (i *Interpreter) Interpret(statements []ast.Statement) {
	for _, stmt := range statements {
		i.execute(stmt)
	}
}

func (i *Interpreter) execute(stmt ast.Statement) {
	stmt.Accept(i)
}

func (i *Interpreter) evaluate(expr ast.Expression) interface{} {
	return expr.Accept(i)
}

// Statements Interpretation
func (i *Interpreter) VisitBlockStmt(stmt *ast.BlockStmt) interface{} {
	i.executeBlock(stmt.Statements, NewEnclosedEnvironment(i.environment))
	return nil
}

func (i *Interpreter) VisitClassStmt(stmt *ast.ClassStmt) interface{} {
	return nil
}

func (i *Interpreter) VisitExpressionStmt(stmt *ast.ExpressionStmt) interface{} {
	i.evaluate(stmt.Expression)
	return nil
}

func (i *Interpreter) VisitFunctionStmt(stmt *ast.FunStmt) interface{} {
	return nil
}

func (i *Interpreter) VisitIfStmt(stmt *ast.IfStmt) interface{} {
	return nil
}

func (i *Interpreter) VisitPrintStmt(stmt *ast.PrintStmt) interface{} {
	value := i.evaluate(stmt.Expression)
	fmt.Println(stringify(value))
	return nil
}

func (i *Interpreter) VisitReturnStmt(stmt *ast.ReturnStmt) interface{} {
	return nil
}

func (i *Interpreter) VisitVarStmt(stmt *ast.VarStmt) interface{} {
	var value interface{}
	if stmt.Initializer != nil {
		value = i.evaluate(stmt.Initializer)
	}
	i.environment.Define(stmt.Name.Lexeme, value)
	return nil
}

func (i *Interpreter) VisitWhileStmt(stmt *ast.WhileStmt) interface{} {
	return nil
}

// Expressions Interpretation
func (i *Interpreter) VisitAssignExpr(expr *ast.Assign) interface{} {
	value := i.evaluate(expr.Value)
	i.environment.Assign(expr.Name, value)
	return value
}

func (i *Interpreter) VisitCallExpr(expr *ast.Call) interface{} {
	return nil
}

func (i *Interpreter) VisitGetExpr(expr *ast.Get) interface{} {
	return nil
}

func (i *Interpreter) VisitGroupingExpr(expr *ast.Grouping) interface{} {
	return nil
}

func (i *Interpreter) VisitLiteralExpr(expr *ast.Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitLogicalExpr(expr *ast.Logical) interface{} {
	return nil
}

func (i *Interpreter) VisitSetExpr(expr *ast.Set) interface{} {
	return nil
}

func (i *Interpreter) VisitSuperExpr(expr *ast.Super) interface{} {
	return nil
}

func (i *Interpreter) VisitThisExpr(expr *ast.This) interface{} {
	return nil
}

func (i *Interpreter) VisitUnaryExpr(expr *ast.Unary) interface{} {
	right := i.evaluate(expr.Right)
	switch expr.Operator.Type {
	case token.MINUS:
		checkNumberOperand(expr.Operator, right)
		if num, ok := right.(float64); ok {
			return -num
		}
	case token.BANG:
		return !isTruthy(right)
	}
	return nil
}

func (i *Interpreter) VisitBinaryExpr(expr *ast.Binary) interface{} {
	leftEval := i.evaluate(expr.Left)
	rightEval := i.evaluate(expr.Right)

	var left float64
	var right float64

	if TypeOf(leftEval) == 'n' && TypeOf(rightEval) == 'n' {
		left, _ = leftEval.(float64)
		right, _ = rightEval.(float64)
	}

	switch expr.Operator.Type {
	case token.PLUS:
		if TypeOf(leftEval) == 'n' && TypeOf(rightEval) == 'n' {
			return left + right
		}
		if TypeOf(leftEval) == 'c' && TypeOf(rightEval) == 'c' {
			left, _ := leftEval.(string)
			right, _ := rightEval.(string)
			return left + right
		}
		panic("Operands must be two numbers or two strings.")
	case token.MINUS:
		checkNumberOperands(expr.Operator, left, right)
		return left - right
	case token.STAR:
		checkNumberOperands(expr.Operator, left, right)
		return left * right
	case token.SLASH:
		// TODO: division by zero.
		checkNumberOperands(expr.Operator, left, right)
		return left / right
	case token.GREATER:
		checkNumberOperand(expr.Operator, left)
		return left > right
	case token.GREATER_EQUAL:
		checkNumberOperands(expr.Operator, left, right)
		return left >= right
	case token.LESS:
		checkNumberOperands(expr.Operator, left, right)
		return left < right
	case token.LESS_EQUAL:
		checkNumberOperands(expr.Operator, left, right)
		return left <= right
	case token.BANG_EQUAL:
		return !isEqual(left, right)
	case token.EQUAL_EQUAL:
		return isEqual(left, right)
	default:
		return nil
	}
}

func (i *Interpreter) VisitVariableExpr(expr *ast.Variable) interface{} {
	return i.environment.Get(expr.Name)
}

// Interpreter helper functions
func (i *Interpreter) executeBlock(statements []ast.Statement, environment *Environment) {
	// save the current env.
	envAct := i.environment

	i.environment = environment
	for _, stmt := range statements {
		i.execute(stmt)
	}

	// restore env
	i.environment = envAct
}

// isTruthy ::= false and nil are Falsey otherwise is Truthy
func isTruthy(object interface{}) bool {
	switch val := object.(type) {
	case bool:
		return val
	case nil:
		return false
	default:
		return true
	}
}

// isEqual()
func isEqual(left interface{}, right interface{}) bool {
	if TypeOf(left) == 'x' && TypeOf(right) == 'x' {
		return true
	}
	if TypeOf(left) == 'x' {
		return false
	}
	return left == right
}

// type
func TypeOf(o interface{}) byte {
	switch o.(type) {
	case string:
		return 'c'
	case float64:
		return 'n'
	case bool:
		return 'l'
	case nil:
		return 'x'
	default:
		return 'u'
	}
}

func checkNumberOperand(operator token.Token, operand interface{}) {
	if _, ok := operand.(float64); ok {
		return
	}
	panic(fmt.Errorf("%v Operand must be a number", operand))
}

func checkNumberOperands(operator token.Token, left interface{}, right interface{}) {
	if _, ok := left.(float64); ok {
		if _, ok := right.(float64); ok {
			return
		}
	}
	panic(fmt.Errorf("%v Operands must be numbers", operator))
}

func stringify(object interface{}) string {
	if TypeOf(object) == 'x' {
		return "nil"
	}
	if TypeOf(object) == 'n' {
		text := fmt.Sprintf("%v", object)
		return text
	}
	return fmt.Sprintf("%v", object)
}
