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

type BaseParser struct {
	Input string
}

func (p *BaseParser) GetLines() []string {
	var lines []string

	reader := strings.NewReader(p.Input)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

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
