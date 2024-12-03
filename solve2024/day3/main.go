package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Token struct {
	Value string
	Type  string
}

type Program struct {
	input   string
	Tokens  []Token
	currIdx int
	peekIdx int
}

// sample of valid expression: mul(123,123)
type Expression struct {
	Mul        *Token
	ParenLeft  *Token
	NumLeft    *Token
	Comma      *Token
	NumRight   *Token
	ParenRight *Token
}

var (
	numericChars  = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	operatorChars = []string{"(", ")", ","}
	keywordChars  = []string{"m", "u", "l"}
)

const (
	NumberToken   = "number"
	KeywordToken  = "keyword"
	OperatorToken = "operator"
)

func main() {
	// f, err := os.ReadFile("./solve2024/inputs/input-day3-test.txt")
	f, err := os.ReadFile("./solve2024/inputs/input-day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	program := Program{
		input:   strings.TrimSpace(string(f)),
		currIdx: 0,
		peekIdx: 1,
		Tokens:  []Token{},
	}
	solvePart1(program)
}

func solvePart1(program Program) {
	program.Parse()

	// program.PrintProgram()
	// fmt.Println()

	expressions := program.ParseExpressions()

	// for _, e := range expressions {
	// 	e.Print()
	// }

	total := 0
	for _, exp := range expressions {
		l, _ := strconv.Atoi(exp.NumLeft.Value)
		r, _ := strconv.Atoi(exp.NumRight.Value)
		total += l * r
	}
	fmt.Println(total)
}

func (p *Program) Parse() {
	for p.currIdx < len(p.input) {
		char := string(p.input[p.currIdx])
		if slices.Contains(numericChars, char) {
			p.parseNumber()
			continue
		}
		if slices.Contains(operatorChars, char) {
			p.parseOperator()
			continue
		}
		if slices.Contains(keywordChars, char) {
			p.parseKeyword()
			continue
		}
		p.currIdx++
		p.peekIdx++
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

func (p *Program) parseOperator() {
	p.peekIdx++
	val := p.input[p.currIdx:p.peekIdx]
	p.Tokens = append(p.Tokens, Token{Value: val, Type: OperatorToken})
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
	exp := Expression{}

	for j := 0; j < len(p.Tokens); j++ {
		tok := p.Tokens[j]
		expected := exp.ExpectedToken()

		if tok.Type != expected.Type || (expected.Value != "" && expected.Value != tok.Value) {
			// fmt.Printf("expected: %v, got: %v\n", expected.Value, tok.Value)
			// fmt.Printf("invalid: ")
			// exp.PrintExpression()

			exp = Expression{}
			continue
		}

		exp.SetToken(&tok)

		if exp.ExpectedToken() == nil {
			// fmt.Printf("valid: ")
			// exp.PrintExpression()

			expressions = append(expressions, exp)
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
		fmt.Printf("%v", tok.Value)
	}
	fmt.Println()
}

func (e *Expression) Print() {
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
	fmt.Println()
}
