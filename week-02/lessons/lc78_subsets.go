package lessons

func subsets(nums []int) [][]int {
	s := &Share{
		set: make([]int, 0),
		ans: make([][]int, 0),
	}

	s.findSubsets(nums, 0)
	return s.ans
}

func (s *Share) findSubsets(nums []int, index int) {
	if index == len(nums) {
		// make a copy
		set := make([]int, len(s.set))
		copy(set, s.set)
		s.ans = append(s.ans, set)
		return
	}

	// 不选
	s.findSubsets(nums, index+1)
	// 选
	s.set = append(s.set, nums[index])
	s.findSubsets(nums, index+1)
	// 还原现场
	s.set = s.set[:len(s.set)-1]
}

type Share struct {
	set []int
	ans [][]int
}
