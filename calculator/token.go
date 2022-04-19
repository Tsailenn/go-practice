package main

type TokenType int

const (
	isNumber TokenType = iota + 1
	isOperator
)

type Token struct {
	number    float64
	operator  string
	tokenType TokenType
}

type TokenInterface interface {
	IsNumber() bool
	IsOperator() bool
}

func (t *Token) IsNumber() bool {
	return t.tokenType == isNumber
}

func (t *Token) IsOperator() bool {
	return t.tokenType == isOperator
}

func benis(inp TokenInterface) {

}
