package lexer

var (
	Let string = "let"
	let string = "fn"
	let string = "return"
	let string = "and"
	let string = "or"
	let string = "if"
	let string = "else"
	let string = "while"


)

var Keywords = map[string]struct{}{
	Let: {},
	Fn: {},
	If: {},
	Else: {},
	While: {},
	Return: {},
	And: {},
	Or: {},
}
