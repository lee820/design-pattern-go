package interpret

import "strings"

//Expression 语句接口
type Expression interface {
	Interpret(context string) bool
}

//TerminalExpression 终端语句类
type TerminalExpression struct {
	Data string
}

//NewTerminalExpression 实例化终端语句类
func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{
		Data: data,
	}
}

//Interpret 终端语句类的解释器
func (te *TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, te.Data) {
		return true
	}
	return false
}

//OrExpression 或语句类
type OrExpression struct {
	Expr1 Expression
	Expr2 Expression
}

//NewOrExpression 实例化或语句
func NewOrExpression(expr1, expr2 Expression) *OrExpression {
	return &OrExpression{
		Expr1: expr1,
		Expr2: expr2,
	}
}

//Interpret 解释器
func (oe *OrExpression) Interpret(context string) bool {
	return oe.Expr1.Interpret(context) || oe.Expr2.Interpret(context)
}

//AndExpression 与语句类
type AndExpression struct {
	Expr1 Expression
	Expr2 Expression
}

//NewAndExpression 实例化与语句
func NewAndExpression(expr1, expr2 Expression) *AndExpression {
	return &AndExpression{
		Expr1: expr1,
		Expr2: expr2,
	}
}

//Interpret 解释器
func (ae *AndExpression) Interpret(context string) bool {
	return ae.Expr1.Interpret(context) && ae.Expr2.Interpret(context)
}
