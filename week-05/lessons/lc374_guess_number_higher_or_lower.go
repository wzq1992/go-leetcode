package lessons

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 0, n

	for l < r {
		mid := (l + r) / 2

		if guess(mid) == 0 {
			return mid
		} else if guess(mid) < 0 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

// leet code 实现
// 这里只是为了编译通过
func guess() int {
	return 0
}
