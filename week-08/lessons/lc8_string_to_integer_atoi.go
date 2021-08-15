package lessons

import "math"

func myAtoi(s string) int {
	n, index := len(s), 0

	// 1. 丢弃前导空格
	for index < n && s[index] == ' ' {
		index++
	}

	// 2. 判断符号
	sign := 1
	if index < n && (s[index] == '+' || s[index] == '-') {
		if s[index] == '-' {
			sign = -1
		}
		index++
	}

	// 3. 处理数字
	val := 0
	for index < n && s[index] >= '0' && s[index] <= '9' {
		if val > (math.MaxInt32-int(s[index]-'0'))/10 {
			if sign == -1 {
				return math.MinInt32
			}

			return math.MaxInt32
		}

		val = val*10 + (int)(s[index]-'0')
		index++
	}

	// 4. 终止条件: 遇到非数字停止
	// 已经包含在数字处理中
	return val * sign
}
