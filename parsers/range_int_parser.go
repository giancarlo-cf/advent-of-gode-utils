package parsers

import (
	"strconv"
	"strings"
)

type RangeIntParser struct {
	BaseParser
}

func (p *RangeIntParser) GetInts(sep string) [][]int {
	var result [][]int

	lines := p.GetLinesBySep(sep)
	for _, line := range lines {
		tokens := strings.Split(line, "-")
		row := make([]int, len(tokens))
		for i, token := range tokens {
			number, err := strconv.Atoi(strings.TrimRight(token, "\n"))
			if err != nil {
				panic(err)
			}
			row[i] = number
		}
		result = append(result, row)
	}

	return result
}
