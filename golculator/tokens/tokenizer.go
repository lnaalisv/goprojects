package tokens

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Tokenize(input string) ([]Token, error) {
	result := []Token{}
	currentToken := ""
	for i, currentChar := range input {
		if isNumber(currentChar) {
			currentToken += string(currentChar)
			if i == len(input)-1 {
				value, _ := strconv.Atoi(currentToken)
				result = append(result, Number{value})
			}
		} else {
			if len(currentToken) > 0 {
				value, _ := strconv.Atoi(currentToken)
				result = append(result, Number{value})
			}

			switch currentToken = string(currentChar); currentToken {
			case "(":
				result = append(result, LeftParenthesis{})
			case ")":
				result = append(result, RightParenthesis{})
			case "+":
				result = append(result, Plus{})
			case "-":
				result = append(result, Minus{})
			default:
				return result, errors.New(fmt.Sprintf("Unknown character %v at position %v", currentToken, i))
			}
			currentToken = ""
		}
	}
	return result, nil
}

func isNumber(input rune) bool {
	numbers := "0123456789"
	return strings.Contains(numbers, string(input))
}
