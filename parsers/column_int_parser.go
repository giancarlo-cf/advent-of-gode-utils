package parsers

import (
	"strconv"
	"strings"
)

type ColumnIntParser struct {
	BaseParser
}

func (p *ColumnIntParser) GetInts(sep string) [][]int {
	var result [][]int

	lines := p.GetLines()
	for _, line := range lines {
		tokens := strings.Split(line, sep)
		for i, token := range tokens {
			if len(result) <= i {
				result = append(result, []int{})
			}
			number, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			result[i] = append(result[i], number)
		}
	}
	return result
}
