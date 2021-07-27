package homework

type TopVotedCandidate struct {
	winners []int
	times   []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	winner := -1
	votes := make(map[int]int, len(persons))
	var wins []int
	for i := range times {
		votes[persons[i]]++
		if winner == -1 {
			winner = persons[i]
		} else {
			if votes[persons[i]] >= votes[winner] {
				winner = persons[i]
			}
		}
		wins = append(wins, winner)
	}

	return TopVotedCandidate{wins, times}
}

func (this *TopVotedCandidate) Q(t int) int {
	l, r := 0, len(this.times)

	for l < r {
		mid := (l + r) / 2

		if this.times[mid] <= t {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return this.winners[l-1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
