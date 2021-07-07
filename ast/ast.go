package ast

import (
	"Glox/token"
	"bytes"
	"fmt"
	"strings"
)

type Node interface {
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Token      token.Token
	Statements []Statement
}

func (p *Program) statementNode() {}
func (p *Program) String() string {
	var out bytes.Buffer
	if len(p.Statements) > 0 {
		for _, stmt := range p.Statements {
			out.WriteString(fmt.Sprintf("%s\n", stmt.String()))
		}
	}
	return out.String()
}

type Block struct {
	Statements []Statement
}

func (b *Block) statementNode() {}
func (b *Block) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	if len(b.Statements) > 0 {
		for _, stmt := range b.Statements {
			out.WriteString(fmt.Sprintf("%s\n", stmt.String()))
		}
	}
	out.WriteString("\n}")
	return out.String()
}

type Class struct {
	Name       token.Token
	Superclass *Variable
	Methods    []Statement
}

func (c *Class) statementNode() {}
func (c *Class) String() string {
	var out bytes.Buffer
	out.WriteString("class ")
	out.WriteString(c.Name.Lexeme)

	if c.Superclass != nil {
		out.WriteString("< ")
		out.WriteString(c.Superclass.String())
	}
	out.WriteString(" {\n")
	if len(c.Methods) > 0 {
		for _, met := range c.Methods {
			out.WriteString(met.String())
		}
	}
	out.WriteString(" }\n")
	return out.String()
}

type ExpressionStmt struct {
	Expression Expression
}

func (e *ExpressionStmt) statementNode() {}
func (e *ExpressionStmt) String() string {
	return e.Expression.String()
}

type Function struct {
	Name   token.Token
	Params []token.Token
	Body   []Statement
}

func (f *Function) statementNode() {}
func (f *Function) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("func %s", f.Name.Lexeme))
	out.WriteString("(")
	if len(f.Params) > 0 {
		var pars []string
		for _, param := range f.Params {
			pars = append(pars, param.Lexeme)
		}
		out.WriteString(fmt.Sprint(strings.Join(pars, ",")))
	}
	out.WriteString("){")
	if f.Body != nil {
		for _, stmt := range f.Body {
			out.WriteString(fmt.Sprintf("%s\n", stmt.String()))
		}
	}
	out.WriteString("\n}")
	return out.String()
}

type IfStmt struct {
	Condition  Expression
	ThenBranch Statement
	ElseBranch Statement
}

func (i *IfStmt) statementNode() {}
func (i *IfStmt) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("if (%s) then", i.Condition.String()))
	out.WriteString(i.ThenBranch.String())

	if i.ElseBranch != nil {
		out.WriteString("else ")
		out.WriteString(i.ElseBranch.String())
	}
	return out.String()
}

type PrintStmt struct {
	Expression Expression
}

func (p *PrintStmt) statementNode() {}
func (p *PrintStmt) String() string {
	return fmt.Sprintf("print(%s)", p.Expression.String())
}

type Return struct {
	Keyword token.Token
	Value   Expression
}

func (r *Return) statementNode() {}
func (r *Return) String() string {
	return fmt.Sprintf("return %s", r.Value.String())
}

type Var struct {
	Name        token.Token
	Initializer Expression
}

func (v *Var) statementNode() {}
func (v *Var) String() string {
	return fmt.Sprintf("var %s = %s", v.Name.Lexeme, v.Initializer.String())
}

type While struct {
	Condition Expression
	Body      Statement
}

func (w *While) statementNode() {}
func (w *While) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("while (%s)%s", w.Condition.String(), w.Body.String()))
	return out.String()
}

// --------------------------------------------------- //
// Expressions derivated Nodes
// --------------------------------------------------- //
type Assign struct {
	Name  token.Token
	Value Expression
}

func (a *Assign) expressionNode() {}
func (a *Assign) String() string {
	return fmt.Sprintf("var %s = %s", a.Name.Lexeme, a.Value.String())
}

type Variable struct {
	Name token.Token
}

func (v *Variable) expressionNode() {}
func (v *Variable) String() string {
	return v.Name.Lexeme
}

type Unary struct {
	Operator token.Token
	Right    Expression
}

func (u *Unary) expressionNode() {}
func (u *Unary) String() string {
	return fmt.Sprintf("(%s %s)", u.Operator.Lexeme, u.Right.String())
}

type Binary struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (b *Binary) expressionNode() {}
func (b *Binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left.String(), b.Operator.Lexeme, b.Right.String())
}

type Call struct {
	Callee    Expression
	Paren     token.Token
	Arguments []Expression
}

func (c *Call) expressionNode() {}
func (c *Call) String() string {
	var out bytes.Buffer
	out.WriteString(c.Callee.String())

	if len(c.Arguments) > 0 {
		var args []string
		for _, arg := range c.Arguments {
			out.WriteString(arg.String())
		}
		out.WriteString(strings.Join(args, ","))
	}

	return out.String()
}

type Get struct {
	Object Expression
	Name   token.Token
}

func (g *Get) expressionNode() {}
func (g *Get) String() string {
	return fmt.Sprintf("get object: %s name: %s>", g.Object.String(), g.Name.Lexeme)
}

type Set struct {
	Object Expression
	Name   token.Token
	Value  Expression
}

func (s *Set) expressionNode() {}
func (s *Set) String() string {
	return fmt.Sprintf("set object: %s name: %s value: %s", s.Object.String(), s.Name.Lexeme, s.Value.String())
}

type Grouping struct {
	Expression Expression
}

func (g *Grouping) expressionNode() {}
func (g *Grouping) String() string {
	return g.Expression.String()
}

type Literal struct {
	Object interface{}
}

func (l *Literal) expressionNode() {}
func (l *Literal) String() string {
	return fmt.Sprintf("%v", l.Object)
}

type Logical struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (l *Logical) expressionNode() {}
func (l *Logical) String() string {
	return fmt.Sprintf("(%s %s %s)", l.Left.String(), l.Operator.Lexeme, l.Right.String())
}

type Super struct {
	Keyword token.Token
	Method  token.Token
}

func (s *Super) expressionNode() {}
func (s *Super) String() string {
	return fmt.Sprintf("super<kw: %s, met: %s>", s.Keyword.Lexeme, s.Method.Lexeme)
}

type This struct {
	Keyword token.Token
}

func (t *This) expressionNode() {}
func (t *This) String() string {
	return t.Keyword.Lexeme
}
