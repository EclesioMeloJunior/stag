package shunting_yard

type operator byte

const (
	add operator = iota
	sub
	mul
	div
)

type Statement interface { 
	isStatement()
}

type Expression interface { 
	isExpression()
}

type Ident struct { 
	varNome string
}
func(Ident) isExpression(){}

type Number struct {
	Value int64
} 
func(Number) isExpression(){}
func(Number) isStatement(){}

type BinaryOperation struct { 
	op operator
	lhs Expression
	rhs Expression
}
func (BinaryOperation) isExpression(){}
func (BinaryOperation) isStatement(){}

type VarAssing struct {
	v string
	value Expression
}
func (VarAssing) isStatement() {}

type Conditional struct { 
	condition Expression
	truePath []Statement
	elsePath []Statement
}
func (Conditional) isStatement() {}

type Loop struct {
	Conditional Expression
	body 		[]Statement
}
func (Loop) isStatement() {}

type FuncDeclaration struct{
	name 		string
	parameters	[]string
	body		[]Statement		
}
func (FuncDeclaration) isStatement() {}

type FuncCall struct{
	name 		string
	arguments	[]Expression
}
func (FuncCall) isExpression() {}