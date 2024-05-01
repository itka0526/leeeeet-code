package bitmanipulation

func reverseBits(num uint32) (ans uint32) {
	ans = 0
	for i := 0; i < 32; i++ {
		checkBit := uint32(1 << i) // i = 5; 0..0100000
		// 0 0 = 0; 0 1 = 0; 1 0 = 0; 1 1 = 1;
		if num&checkBit == 0 {
			ans <<= 1
		} else {
			ans = (ans << 1) | 1
		}
	}
	return ans
}
