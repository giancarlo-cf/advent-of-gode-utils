package parsers

type MatrixRuneParser struct {
	BaseParser
}

func (p *MatrixRuneParser) GetStrings() [][]rune {
	var result [][]rune

	lines := p.GetLines()
	for _, line := range lines {
		runes := []rune(line)
		result = append(result, runes)
	}

	return result
}
