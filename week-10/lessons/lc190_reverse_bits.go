package lessons

func reverseBits(num uint32) uint32 {
	var ans uint32

	for i := uint32(0); i < 32; i++ {
		ans = (ans << 1) | (num >> i & 1)
	}

	return ans
}
