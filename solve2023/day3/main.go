package main

import (
	"fmt"
	"log"
	"os"
	"slices"
)

type Token struct {
	Value string
	Type  string
}

type Expression struct {
	Tokens []Token
}

type Program struct {
	input       string
	Tokens      []Token
	Expressions []Expression
	currIdx     int
	peekIdx     int
}

const (
	NumberType = "NUMBER"
	SymbolType = "SYMBOL"
)

func main() {
	// f, err := os.ReadFile("./solve2023/inputs/input-day3.txt")
	f, err := os.ReadFile("./solve2023/inputs/input-day3-test.txt")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	part1(string(f))
	part2(string(f))
}

func part1(input string) {
	p := Program{input: input, currIdx: 0, peekIdx: 1}

	p.ScanTokens()
	p.Print()

	p.Parse()
}

func part2(input string) {
}

func (p *Program) ScanTokens() {
	for p.currIdx < len(p.input) {
		char := string(p.input[p.currIdx])
		if isNumber(char) {
			p.scanNumber()
		} else if isSymbol(char) {
			p.scanSymbol()
		}
		p.currIdx++
		p.peekIdx++
	}
}

// TODO:
func (p *Program) Parse() {
	for i := 0; i < len(p.Tokens); i++ {
		tok := p.Tokens[i]
		if tok.Type == NumberType {
			// p.Expressions = append(p.Expressions, Expression{Tokens: []Token{tok}})
		}
	}
}

func (p *Program) peek() string {
	if p.peekIdx < len(p.input) {
		return string(p.input[p.peekIdx])
	}
	return ""
}

func (p *Program) scanNumber() {
	for isNumber(p.peek()) {
		p.peekIdx++
	}
	tok := Token{Value: p.input[p.currIdx:p.peekIdx], Type: NumberType}
	p.Tokens = append(p.Tokens, tok)
	p.currIdx = p.peekIdx
}

func (p *Program) scanSymbol() {
	for isSymbol(p.peek()) {
		p.peekIdx++
	}
	tok := Token{Value: p.input[p.currIdx:p.peekIdx], Type: SymbolType}
	p.Tokens = append(p.Tokens, tok)
	p.currIdx = p.peekIdx
}

func (p Program) Print() {
	for _, tok := range p.Tokens {
		fmt.Printf("%s[%s] ", tok.Type, tok.Value)
	}
	fmt.Println()
}

func isNumber(c string) bool {
	return slices.Contains([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, c)
}

func isSymbol(c string) bool {
	return slices.Contains([]string{"+", "-", "*", "/", "@", "#", "$", "%", "^", "&", "=", "~", "`", "?", "!", "<", ">", ";", ":", "[", "]", "{", "}", "|", "\\"}, c)
}
