package parse

import (
	"fmt"
	"strconv"

	"github.com/BriceLucifer/golisp/lexer"
)

type Expr interface{}

type Symbol string

type Number float64

type List []Expr

type Parser struct {
	lexer *lexer.Lexer
	token lexer.Token
}

func NewParser(input string) *Parser {
	lexer := lexer.NewLexer(input)
	return &Parser{lexer: lexer, token: lexer.NextToken()}
}

func (p *Parser) Parse() Expr {
	return p.parseExpr()
}

func (p *Parser) parseExpr() Expr {
	switch p.token.Type {
	case lexer.TokenNumber:
		n, _ := strconv.ParseFloat(p.token.Value, 64)
		p.nextToken()
		return Number(n)
	case lexer.TokenSymbol:
		s := Symbol(p.token.Value)
		p.nextToken()
		return s
	case lexer.TokenLParen:
		p.nextToken()
		return p.parseList()
	default:
		panic(fmt.Sprintf("unexpected token: %v", p.token))
	}
}

func (p *Parser) parseList() List {
	var list List
	for p.token.Type != lexer.TokenRParen {
		list = append(list, p.parseExpr())
	}
	p.nextToken() // Skip the closing parenthesis
	return list
}

func (p *Parser) nextToken() {
	p.token = p.lexer.NextToken()
}
