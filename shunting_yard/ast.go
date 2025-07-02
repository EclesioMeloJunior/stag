package shunting_yard

type Operation byte

const (
	Add Operation = iota
	Sub
	Mul
	Div
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

func (Ident) isExpression() {}

type Number struct {
	Value int64
}

func (*Number) isExpression() {}
func (*Number) isStatement()  {}

type BinaryOperation struct {
	Op  Operation
	Lhs Expression
	Rhs Expression
}

func (*BinaryOperation) isExpression() {}
func (*BinaryOperation) isStatement()  {}

type VarAssing struct {
	v     string
	value Expression
}

func (VarAssing) isStatement() {}

type Conditional struct {
	condition Expression
	truePath  []Statement
	elsePath  []Statement
}

// criar uma nova branch para essas funções
func (Conditional) isStatement() {}

type Loop struct{}

func (Loop) isStatement() {}

type FuncDeclaration struct{}

func (FuncDeclaration) isStatement() {}

type FuncCall struct{}

func (FuncCall) isExpression() {}
