package lexer

import (
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenNumber
	TokenSymbol
	TokenLParen
	TokenRParen
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input *strings.Reader
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: strings.NewReader(input)}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()
	ch := l.peek()
	switch {
	case ch == 0:
		return Token{Type: TokenEOF}
	case ch == '(':
		l.next()
		return Token{Type: TokenLParen}
	case ch == ')':
		l.next()
		return Token{Type: TokenRParen}
	case unicode.IsDigit(ch):
		return Token{Type: TokenNumber, Value: l.readNumber()}
	case unicode.IsLetter(ch) || ch == '-':
		return Token{Type: TokenSymbol, Value: l.readSymbol()}
	default:
		l.next() // Consume and ignore unknown characters
		return l.NextToken()
	}
}

func (l *Lexer) peek() rune {
	ch, _, err := l.input.ReadRune()
	if err != nil {
		return 0
	}
	l.input.UnreadRune()
	return ch
}

func (l *Lexer) next() rune {
	ch, _, err := l.input.ReadRune()
	if err != nil {
		return 0
	}
	return ch
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.peek()) {
		l.next()
	}
}

func (l *Lexer) readNumber() string {
	var sb strings.Builder
	for unicode.IsDigit(l.peek()) || l.peek() == '.' {
		sb.WriteRune(l.next())
	}
	return sb.String()
}

func (l *Lexer) readSymbol() string {
	var sb strings.Builder
	for unicode.IsLetter(l.peek()) || unicode.IsDigit(l.peek()) || l.peek() == '-' {
		sb.WriteRune(l.next())
	}
	return sb.String()
}