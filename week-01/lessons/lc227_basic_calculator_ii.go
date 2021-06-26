package lessons
//
// import (
// 	"fmt"
// 	"strconv"
// )
//
// func calculate(s string) int {
// 	var tokens []string
// 	var ops []byte
// 	var val int
// 	var numStart bool
// 	needZero := true
//
// 	for i := 0; i < len(s); i++ {
// 		ch := s[i]
// 		if ch >= '0' && ch <= '9' {
// 			val = val*10 + int(ch-'0')
// 			numStart = true
// 			continue
// 		} else if numStart {
// 			tokens = append(tokens, strconv.Itoa(val))
// 			val = 0
// 			numStart = false
// 			needZero = false
// 		}
//
// 		if ch == ' ' {
// 			continue
// 		}
//
// 		if ch == '(' {
// 			ops = append(ops, ch)
// 			needZero = true
// 			continue
// 		}
//
// 		if ch == ')' {
// 			for ops[len(ops)-1] != '(' {
// 				tokens = append(tokens, string(ops[len(ops)-1]))
// 				ops = ops[:len(ops)-1]
// 			}
// 			ops = ops[:len(ops)-1]
// 			needZero = false
// 			continue
// 		}
//
// 		if needZero {
// 			tokens = append(tokens, "0")
// 		}
//
// 		for len(ops) != 0 && getRank(ops[len(ops)-1]) >= getRank(ch) {
// 			tokens = append(tokens, string(ops[len(ops)-1]))
// 			ops = ops[:len(ops)-1]
// 		}
//
// 		ops = append(ops, ch)
// 		needZero = true
// 	}
//
// 	if numStart {
// 		tokens = append(tokens, strconv.Itoa(val))
// 	}
//
// 	for len(ops) != 0 {
// 		tokens = append(tokens, string(ops[len(ops)-1]))
// 		ops = ops[:len(ops)-1]
// 	}
//
// 	fmt.Println(tokens)
// 	return evalRPN(tokens)
// }
//
// func getRank(ch byte) int {
// 	switch ch {
// 	case '+', '-':
// 		return 1
// 	case '*', '/':
// 		return 2
// 	default:
// 		return 0
// 	}
// }
//
// func evalRPN(tokens []string) int {
// 	var stack []int
//
// 	for _, s := range tokens {
// 		if s == "+" || s == "-" || s == "*" || s == "/" {
// 			v1 := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			v2 := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
//
// 			stack = append(stack, calc(s, v2, v1))
// 		} else {
// 			v, _ := strconv.Atoi(s)
// 			stack = append(stack, v)
// 		}
// 	}
//
// 	return stack[len(stack)-1]
// }
//
// func calc(op string, v1, v2 int) int {
// 	switch op {
// 	case "+":
// 		return v1 + v2
// 	case "-":
// 		return v1 - v2
// 	case "*":
// 		return v1 * v2
// 	case "/":
// 		return v1 / v2
// 	}
//
// 	return 0
// }
