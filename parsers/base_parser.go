package parsers

import (
	"bufio"
	"strings"
)

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

func (p *BaseParser) GetLinesBySep(sep string) []string {
	return strings.Split(p.Input, sep)
}
