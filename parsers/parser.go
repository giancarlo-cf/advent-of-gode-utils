package parsers

type Parser interface {
	GetLines() []string
}

type IntParser interface {
	Parser
	GetInts(sep string) [][]int
}

type StringParser interface {
	Parser
	GetStrings(sep string) [][]string
}
