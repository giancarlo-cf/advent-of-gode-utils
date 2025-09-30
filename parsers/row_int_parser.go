package parsers

import (
	"strconv"
	"strings"
)

type RowIntParser struct {
	BaseParser
}

func (p *RowIntParser) GetInts(sep string) [][]int {
	var result [][]int

	lines := p.GetLines()
	for _, line := range lines {
		tokens := strings.Split(line, sep)
		row := make([]int, len(tokens), len(tokens))
		for i, token := range tokens {
			number, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			row[i] = number
		}
		result = append(result, row)
	}

	return result
}
