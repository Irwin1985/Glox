package ast

import (
	"bytes"
	"fmt"
	"strings"
)

func Beautify(node interface{}) string {
	var out bytes.Buffer
	switch node := node.(type) {
	case *BlockStmt:
		out.WriteString("{\n")
		if len(node.Statements) > 0 {
			for _, stmt := range node.Statements {
				out.WriteString(fmt.Sprintf("%s\n", stmt.String()))
			}
		}
		out.WriteString("\n}")
	case *ClassStmt:
		out.WriteString("class ")
		out.WriteString(node.Name.Lexeme)

		if node.Superclass != nil {
			out.WriteString("< ")
			out.WriteString(node.Superclass.String())
		}
		out.WriteString(" {\n")
		if len(node.Methods) > 0 {
			for _, met := range node.Methods {
				out.WriteString(met.String())
			}
		}
		out.WriteString(" }\n")
	case *ExpressionStmt:
		out.WriteString(node.Expression.String())
	case *FunStmt:
		out.WriteString(fmt.Sprintf("func %s", node.Name.Lexeme))
		out.WriteString("(")
		if len(node.Params) > 0 {
			var pars []string
			for _, param := range node.Params {
				pars = append(pars, param.Lexeme)
			}
			out.WriteString(fmt.Sprint(strings.Join(pars, ",")))
		}
		out.WriteString("){")
		if node.Body != nil {
			for _, stmt := range node.Body {
				out.WriteString(fmt.Sprintf("%s\n", stmt.String()))
			}
		}
		out.WriteString("\n}")
	case *IfStmt:
		out.WriteString(fmt.Sprintf("if (%s) then", node.Condition.String()))
		out.WriteString(node.ThenBranch.String())

		if node.ElseBranch != nil {
			out.WriteString("else ")
			out.WriteString(node.ElseBranch.String())
		}
	case *PrintStmt:
		out.WriteString(fmt.Sprintf("print(%s)", node.Expression.String()))
	case *ReturnStmt:
		out.WriteString(fmt.Sprintf("return %s", node.Value.String()))
	case *VarStmt:
		out.WriteString(fmt.Sprintf("var %s = %s", node.Name.Lexeme, node.Initializer.String()))
	case *WhileStmt:
		out.WriteString(fmt.Sprintf("while (%s)%s", node.Condition.String(), node.Body.String()))
	case *Assign:
		out.WriteString(fmt.Sprintf("var %s = %s", node.Name.Lexeme, node.Value.String()))
	case *Variable:
		out.WriteString(node.Name.Lexeme)
	case *Unary:
		out.WriteString(fmt.Sprintf("(%s %s)", node.Operator.Lexeme, node.Right.String()))
	case *Binary:
		out.WriteString(fmt.Sprintf("(%s %s %s)", node.Left.String(), node.Operator.Lexeme, node.Right.String()))
	case *Call:
		out.WriteString(node.Callee.String())
		if len(node.Arguments) > 0 {
			var args []string
			for _, arg := range node.Arguments {
				out.WriteString(arg.String())
			}
			out.WriteString(strings.Join(args, ","))
		}
	case *Get:
		out.WriteString(fmt.Sprintf("get object: %s name: %s>", node.Object.String(), node.Name.Lexeme))
	case *Set:
		out.WriteString(fmt.Sprintf("set object: %s name: %s value: %s", node.Object.String(), node.Name.Lexeme, node.Value.String()))
	case *Grouping:
		out.WriteString(node.Expression.String())
	case *Literal:
		out.WriteString(fmt.Sprintf("%v", node.Value))
	case *Logical:
		out.WriteString(fmt.Sprintf("(%s %s %s)", node.Left.String(), node.Operator.Lexeme, node.Right.String()))
	case *Super:
		out.WriteString(fmt.Sprintf("super<kw: %s, met: %s>", node.Keyword.Lexeme, node.Method.Lexeme))
	case *This:
		out.WriteString(node.Keyword.Lexeme)
	}
	return out.String()
}
