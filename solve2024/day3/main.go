package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type Token struct {
	Value string
	Type  string
}

type Program struct {
	input    string
	Tokens   []Token
	Extended bool
	currIdx  int
	peekIdx  int
}

// sample of valid expression: mul(123,123)
type Expression struct {
	// Enabled    *Token
	Mul        *Token
	ParenLeft  *Token
	NumLeft    *Token
	Comma      *Token
	NumRight   *Token
	ParenRight *Token
}

var (
	numericChars   = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	keywordChars   = []string{"m", "u", "l"}
	extendedChars  = []string{"d", "o", "n", "'", "t", "(", ")"}
	extendedTokens = []string{"do()", "don't()"}
)

const (
	NumberToken   = "number"
	KeywordToken  = "keyword"
	OperatorToken = "operator"
)

func main() {
	{
		// f, err := os.ReadFile("./solve2024/inputs/input-day3-test.txt")
		f, err := os.ReadFile("./solve2024/inputs/input-day3.txt")
		if err != nil {
			log.Fatal(err)
		}
		solvePart1(string(f))
	}

	{
		// f, err := os.ReadFile("./solve2024/inputs/input-day3-part2-test.txt")
		f, err := os.ReadFile("./solve2024/inputs/input-day3.txt")
		if err != nil {
			log.Fatal(err)
		}
		solvePart2(string(f))
	}
}

func solvePart1(input string) {
	// fmt.Println("input:", input)
	program := Program{
		input:    string(input),
		currIdx:  0,
		peekIdx:  0,
		Tokens:   []Token{},
		Extended: false,
	}
	program.Parse()
	expressions := program.ParseExpressions()
	fmt.Println("part1:", calcTotal(expressions))
}

func solvePart2(input string) {
	program := Program{
		input:    string(input),
		currIdx:  0,
		peekIdx:  0,
		Tokens:   []Token{},
		Extended: true,
	}

	program.Parse()

	// fmt.Printf("\ntokens:\n")
	// program.Print()
	// fmt.Println()

	expressions := program.ParseExpressions()
	// {
	// fmt.Printf("\nexpressions:\n")
	// for _, e := range expressions {
	// 	e.Print()
	// 	fmt.Printf(" ")
	// }
	// }
	fmt.Println("part2:", calcTotal(expressions))
}

func (p *Program) Parse() {
	for p.currIdx < len(p.input) {
		char := string(p.input[p.currIdx])

		if p.Extended && char == "d" {
			p.parseExtended()
			continue
		}
		if slices.Contains(numericChars, char) {
			p.parseNumber()
			continue
		}
		if slices.Contains(keywordChars, char) {
			p.parseKeyword()
			continue
		}

		p.parseSpecials()
	}
}

func (p *Program) parseNumber() {
	for slices.Contains(numericChars, p.peek()) {
		p.peekIdx++
	}
	if val := p.input[p.currIdx:p.peekIdx]; len(val) > 0 && len(val) <= 3 {
		p.Tokens = append(p.Tokens, Token{Value: val, Type: NumberToken})
	}
	p.currIdx = p.peekIdx
}

func (p *Program) parseKeyword() {
	for slices.Contains(keywordChars, p.peek()) {
		p.peekIdx++
	}
	if val := p.input[p.currIdx:p.peekIdx]; val == "mul" {
		p.Tokens = append(p.Tokens, Token{Value: val, Type: KeywordToken})
	}
	p.currIdx = p.peekIdx
}

func (p *Program) parseSpecials() {
	p.peekIdx++
	val := p.input[p.currIdx:p.peekIdx]
	p.Tokens = append(p.Tokens, Token{Value: val, Type: OperatorToken})
	p.currIdx = p.peekIdx
}

func (p *Program) parseExtended() {
	for slices.Contains(extendedChars, p.peek()) {
		if p.peek() == ")" {
			p.peekIdx++
			break
		} else {
			p.peekIdx++
		}
	}
	if val := p.input[p.currIdx:p.peekIdx]; slices.Contains(extendedTokens, val) {
		p.Tokens = append(p.Tokens, Token{Value: val, Type: KeywordToken})
	}
	p.currIdx = p.peekIdx
}

func (p *Program) peek() string {
	if p.peekIdx > len(p.input)-1 {
		return ""
	}
	return string(p.input[p.peekIdx])
}

func (p *Program) ParseExpressions() []Expression {
	expressions := []Expression{}
	skip := false
	exp := Expression{}

	for j := 0; j < len(p.Tokens); j++ {
		tok := p.Tokens[j]

		if p.Extended && slices.Contains(extendedTokens, tok.Value) {
			skip = tok.Value == "don't()"
			// expressions = append(expressions, Expression{Enabled: &tok})
			continue
		}

		expected := exp.ExpectedToken()
		if tok.Type != expected.Type || (expected.Value != "" && expected.Value != tok.Value) {
			exp = Expression{}
			continue
		}

		exp.SetToken(&tok)

		if exp.ExpectedToken() == nil {
			if !skip {
				expressions = append(expressions, exp)
			}
			exp = Expression{}
		}
	}

	return expressions
}

func (e *Expression) ExpectedToken() *Token {
	if e.Mul == nil {
		return &Token{Value: "mul", Type: KeywordToken}
	}
	if e.ParenLeft == nil {
		return &Token{Value: "(", Type: OperatorToken}
	}
	if e.NumLeft == nil {
		return &Token{Type: NumberToken, Value: ""}
	}
	if e.Comma == nil {
		return &Token{Value: ",", Type: OperatorToken}
	}
	if e.NumRight == nil {
		return &Token{Type: NumberToken, Value: ""}
	}
	if e.ParenRight == nil {
		return &Token{Value: ")", Type: OperatorToken}
	}
	return nil
}

func (e *Expression) SetToken(tok *Token) {
	switch tok.Type {
	case KeywordToken:
		e.Mul = tok

	case OperatorToken:
		switch tok.Value {
		case "(":
			e.ParenLeft = tok
		case ",":
			e.Comma = tok
		case ")":
			e.ParenRight = tok
		}

	case NumberToken:
		if e.NumLeft == nil {
			e.NumLeft = tok
		} else {
			e.NumRight = tok
		}
	}
}

func (p *Program) Print() {
	for _, tok := range p.Tokens {
		fmt.Printf("%v ", tok.Value)
	}
	fmt.Println()
}

func (e *Expression) Print() {
	// if e.Enabled != nil {
	//	fmt.Printf("%v", e.Enabled.Value)
	//	return
	// }
	if e.Mul != nil {
		fmt.Printf("%v", e.Mul.Value)
	}
	if e.ParenLeft != nil {
		fmt.Printf("%v", e.ParenLeft.Value)
	}
	if e.NumLeft != nil {
		fmt.Printf("%v", e.NumLeft.Value)
	}
	if e.Comma != nil {
		fmt.Printf("%v", e.Comma.Value)
	}
	if e.NumRight != nil {
		fmt.Printf("%v", e.NumRight.Value)
	}
	if e.ParenRight != nil {
		fmt.Printf("%v", e.ParenRight.Value)
	}
}

func calcTotal(expressions []Expression) int {
	total := 0
	for _, exp := range expressions {
		if exp.Enabled != nil {
			continue
		}

		l, _ := strconv.Atoi(exp.NumLeft.Value)
		r, _ := strconv.Atoi(exp.NumRight.Value)
		total += l * r
	}
	return total
}
