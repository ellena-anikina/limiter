package pwdvalidator

import (
	"unicode"
	"unicode/utf8"
)

type validator func(s string) bool

func digits(s string) bool {
	var result bool = false
	for _, r := range s {
		if unicode.IsDigit(r) {
			result = true
		}
	}
	return result
}

func letters(s string) bool {
	var result bool = false
	for _, r := range s {
		if unicode.IsLetter(r) {
			result = true
		}
	}
	return result
}

func minlen(length int) validator {
	return func(s string) bool {
		return utf8.RuneCountInString(s) >= length
	}
}

func and(funcs ...validator) validator {
	return func(s string) bool {
		result := true
		for _, fn := range funcs {
			if !fn(s) {
				result = false
			}
		}
		return result
	}
}

func or(funcs ...validator) validator {
	// ...
	return func(s string) bool {
		result := false
		for _, fn := range funcs {
			if fn(s) {
				result = true
			}
		}
		return result
	}
}

type password struct {
	value string
	validator
}

func (p *password) isValid() bool {
	return p.validator(p.value)
}
