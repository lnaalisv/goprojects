package tokens

import "strconv"

type Token interface {
	isToken()
}

func (_ LeftParenthesis) isToken() {}
func (_ LeftParenthesis) String() string {
	return "("
}

type LeftParenthesis struct{}

func (_ RightParenthesis) isToken() {}
func (_ RightParenthesis) String() string {
	return ")"
}

type RightParenthesis struct{}

func (_ Plus) isToken() {}
func (_ Plus) String() string {
	return "+"
}

type Plus struct{}

func (_ Minus) isToken() {}
func (_ Minus) String() string {
	return "-"
}

type Minus struct{}

func (_ Number) isToken() {}
func (n Number) String() string {
	return strconv.Itoa(n.value)
}

type Number struct {
	value int
}
