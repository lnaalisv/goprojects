package tokens

import "testing"

func areEqual(a []Token, b []Token) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTokenize(t *testing.T) {
	result, _ := Tokenize("+0123-500()")
	expected := []Token{
		Plus{}, Number{123}, Minus{}, Number{500}, LeftParenthesis{}, RightParenthesis{},
	}
	if !areEqual(result, expected) {
		t.Error("Expected", expected, "Got", result)
	}
}
