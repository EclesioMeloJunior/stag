package lexer

var (
	Let string = "let"
	Fn  string = "fn"
	Return string = "return"
	And   string = "and"
	Or    string = "or"
	If    string = "if"
	Else  string = "else"
	While string = "while"


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
