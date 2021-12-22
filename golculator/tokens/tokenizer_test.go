package tokens

import "testing"

func TestTokenize(t *testing.T) {
	result, _ := Tokenize("+0123-500()")
	if result != nil {
		t.Error("Expected foo received", result)
	}
}
