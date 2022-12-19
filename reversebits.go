package piscine

func ReverseBits(oct byte) byte {
	result := byte(0)
	for i := 0; i < 8; i++ {
		oct = (oct >> 1) & 1
		result = result | (oct << (7 - i))
	}
	return result
}
