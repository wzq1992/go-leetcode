package lessons

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int

	for _, s := range tokens {
		if s == "+" || s == "-" || s == "*" || s == "/" {
			v1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			v2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			stack = append(stack, calc(s, v2, v1))
		} else {
			v, _ := strconv.Atoi(s)
			stack = append(stack, v)
		}
	}

	return stack[len(stack)-1]
}

func calc(op string, v1, v2 int) int {
	switch op {
	case "+":
		return v1 + v2
	case "-":
		return v1 - v2
	case "*":
		return v1 * v2
	case "/":
		return v1 / v2
	}

	return 0
}
