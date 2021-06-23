package homework

func plusOne(digits []int) []int {
	var res []int

	// 添加一个最高位，因为有可能 +1 后发生连续进位导致数字位数 +1
	// case: 9999 +1 ---> 10000
	res = append(res, 0)
	res = append(res, digits...)

	// 从最末位开始计算
	// +1 后未发生进位, 直接当前未数字 +1 并结束循环
	// +1 后发生进位，当前位数字置为 0 继续处理前一位
	for i := len(digits); i >= 0; i-- {
		if res[i] < 9 {
			res[i] += 1
			break
		}

		res[i] = 0
	}

	// 最高位为 0 时为无效位，返回时去除
	if res[0] == 0 {
		return res[1:]
	}

	return res
}
