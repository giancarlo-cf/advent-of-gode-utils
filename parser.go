package advent_of_gode_utils

import (
	"bufio"
	"strconv"
	"strings"
)

type Parser interface {
	GetLines() []string
}

type IntParser interface {
	Parser
	GetInts(sep string) [][]int
}

type ColumnIntParser struct {
	Input string
}

func (p *ColumnIntParser) GetLines() []string {
	var lines []string

	reader := strings.NewReader(p.Input)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
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
