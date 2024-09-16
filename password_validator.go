// Напишите программу, которая проверяет корректность пароля.

// Следуйте указаниям по тексту программы.
// Чтобы протестировать своё решение, запустите go test -v в папке с упражнением.
// Не меняйте сигнатуры функций, определение типа password, чтобы тесты прошли корректно.

package pwdvalidator

import (
	"unicode"
	"unicode/utf8"
)

// validator проверяет строку на соответствие некоторому условию
// и возвращает результат проверки
type validator func(s string) bool

// digits возвращает true, если s содержит хотя бы одну цифру
// согласно unicode.IsDigit(), иначе false
func digits(s string) bool {
	var result bool = false
	for _, r := range s {
		if unicode.IsDigit(r) {
			result = true
		}
	}
	return result
}

// letters возвращает true, если s содержит хотя бы одну букву
// согласно unicode.IsLetter(), иначе false
func letters(s string) bool {
	var result bool = false
	for _, r := range s {
		if unicode.IsLetter(r) {
			result = true
		}
	}
	return result
}

// minlen возвращает валидатор, который проверяет, что длина
// строки согласно utf8.RuneCountInString() - не меньше указанной
func minlen(length int) validator {
	return func(s string) bool {
		return utf8.RuneCountInString(s) >= length
	}
}

// and возвращает валидатор, который проверяет, что все
// переданные ему валидаторы вернули true
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

// or возвращает валидатор, который проверяет, что хотя бы один
// переданный ему валидатор вернул true
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

// password содержит строку со значением пароля и валидатор
type password struct {
	value string
	validator
}

// isValid() проверяет, что пароль корректный, согласно
// заданному для пароля валидатору
func (p *password) isValid() bool {
	// ...
	return p.validator(p.value)
}
