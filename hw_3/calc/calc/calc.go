package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	errBadExpression  = errors.New("bad expression")
	errDivisionByZero = errors.New("division by zero")
	errBadNumber      = errors.New("bad number")
)

// Calc вычисляет значение выражения, заданного строкой.
// Возвращает результат вычисления или ошибку, если выражение некорректно.
// Пример: "1.23 + 54 * (6/7 * (9 - 8))" -> 47.51571428571428, nil
func Calc(expression string) (float64, error) {
	rpn := toRPN(tokenizeExpression(expression))
	res, err := calc(rpn)
	if err != nil {
		return 0, fmt.Errorf("calc fail: %v", err)
	}
	return res, nil
}

func calc(rpn []string) (float64, error) {
	stack := make([]float64, 0, len(rpn))
	var curResult float64

	for _, token := range rpn {
		switch token {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, errBadExpression
			}
			rhs := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			lhs := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch token {
			case "+":
				curResult = lhs + rhs
			case "-":
				curResult = lhs - rhs
			case "*":
				curResult = lhs * rhs
			default:
				if rhs == 0 {
					return 0, errDivisionByZero
				}
				curResult = lhs / rhs
			}
			stack = append(stack, curResult)
		default:
			value, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, errBadNumber
			}
			stack = append(stack, value)
		}
	}

	if len(stack) != 1 {
		return 0, errBadExpression
	}

	return stack[0], nil
}

// tokenizeExpression разбивает строку выражения на отдельные токены.
// Токены могут быть числами, операторами или скобками.
// Пример: "1.23 + 54 * (6/7 * (9 - 8))" ->
// []string{"1.23", "+", "54", "*", "(", "6", "/", "7", "*", "(", "9", "-", "8", ")", ")"}
func tokenizeExpression(expression string) []string {
	tokens := make([]string, 0, utf8.RuneCountInString(expression))
	number := make([]string, 0, utf8.RuneCountInString(expression))
	expression = strings.Replace(expression, " ", "", -1)

	for _, char := range strings.Split(expression, "") {
		if ((char >= "0") && (char <= "9")) || char == "." || char == "," {
			number = append(number, char)
			continue
		}
		if len(number) > 0 {
			tokens = append(tokens, strings.Join(number, ""))
			number = make([]string, 0, len(expression))
		}
		tokens = append(tokens, char)
	}
	if len(number) > 0 {
		tokens = append(tokens, strings.Join(number, ""))
	}
	return tokens
}

func operatorPriority(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

//

// toRPN преобразует инфиксное выражение в обратную польскую нотацию (ОПН).
// Пример: []string{"1.23", "+", "54", "*", "(", "6", "/", "7", "*", "(", "9", "-", "8", ")", ")"} ->
// []string{"1.23", "54", "6", "7", "/", "9", "8", "-", "*", "*", "+"}
func toRPN(tokens []string) []string {
	stack := make([]string, 0, len(tokens))
	rpnTokens := make([]string, 0, len(tokens))
	for _, token := range tokens {
		switch {
		case token >= "0" && string([]rune(token)[0]) <= "9":
			rpnTokens = append(rpnTokens, token)
		case token == "(":
			stack = append(stack, token)
		case token == ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				rpnTokens = append(rpnTokens, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			for len(stack) > 0 && operatorPriority(stack[len(stack)-1]) >= operatorPriority(token) {
				rpnTokens = append(rpnTokens, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		}
	}

	for len(stack) > 0 {
		rpnTokens = append(rpnTokens, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return rpnTokens
}
