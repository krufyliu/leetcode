package BasicCalculatorII

import "strconv"

func calculate(s string) int {
	l := len(s)
	tokens := make([]string, 0, 4)
	for i := 0; i < l; i++ {
		if s[i] == ' ' {
			continue
		} else if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			tokens = append(tokens, s[i:i+1])
			continue
		}
		j := i
		for j < l {
			if s[j] >= '0' && s[j] <= '9' {
				j++
				continue
			}
			break
		}
		tokens = append(tokens, s[i:j])
		i = j - 1
	}

	var result, left, right int
	var stack = make([]string, 0, len(tokens))

	var i = 0
	// 先处理乘除法
	for i < len(tokens) {
		if tokens[i] != "*" && tokens[i] != "/" {
			stack = append(stack, tokens[i])
			i++
			continue
		}
		op := tokens[i]
		left, _ = strconv.Atoi(stack[len(stack)-1])
		stack = stack[0 : len(stack)-1]
		i++
		right, _ = strconv.Atoi(tokens[i])
		if op == "*" {
			stack = append(stack, strconv.Itoa(left*right))
		} else {
			stack = append(stack, strconv.Itoa(left/right))
		}
		i++
	}
	// 处理加减法
	result, _ = strconv.Atoi(stack[0])
	i = 1
	for i < len(stack) {
		token := stack[i]
		if token == "+" {
			i++
			right, _ = strconv.Atoi(stack[i])
			result += right
		} else if token == "-" {
			i++
			right, _ = strconv.Atoi(stack[i])
			result -= right
		}
		i++
	}

	return result
}
