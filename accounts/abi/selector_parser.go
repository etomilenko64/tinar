package abi

import (
	"fmt"
)

type SelectorMarshaling struct {
	Name   string               `json:"name"`
	Type   string               `json:"type"`
	Inputs []ArgumentMarshaling `json:"inputs"`
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isIdentifierSymbol(c byte) bool {
	return c == '$' || c == '_'
}

func parseToken(unescapedSelector string, isIdent bool) (string, string, error) {
	if len(unescapedSelector) == 0 {
		return "", "", fmt.Errorf("empty token")
	}
	firstChar := unescapedSelector[0]
	position := 1
	if !(isAlpha(firstChar) || (isIdent && isIdentifierSymbol(firstChar))) {
		return "", "", fmt.Errorf("invalid token start: %c", firstChar)
	}
	for position < len(unescapedSelector) {
		char := unescapedSelector[position]
		if !(isAlpha(char) || isDigit(char) || (isIdent && isIdentifierSymbol(char))) {
			break
		}
