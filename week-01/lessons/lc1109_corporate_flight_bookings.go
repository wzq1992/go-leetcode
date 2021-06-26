package lessons

func corpFlightBookings(bookings [][]int, n int) []int {
	delta := make([]int, n+2)

	for _, book := range bookings {
		first := book[0]
		last := book[1]
		seats := book[2]

		delta[first] += seats
		delta[last+1] -= seats
	}

	res := make([]int, n+1)

	for i := 1; i <= n; i++ {
		res[i] = res[i-1] + delta[i]
	}

	return res[1:]
}
