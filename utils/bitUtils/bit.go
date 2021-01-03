package bitUtils

func SelectBits(cur byte, highBit, lowBit byte) byte {
	mask := byte(0)

	if highBit > 7 || lowBit > 7 {
		return 0
	}

	for i := lowBit; i <= highBit; i++ {
		mask |= 1 << i
	}

	return cur & mask >> lowBit
}

func SetBits(cur *byte, highBit, lowBit, newVal byte) {
	mask := byte(0)

	if highBit > 7 || lowBit > 7 {
		return
	}

	for i := lowBit; i <= highBit; i++ {
		mask |= 1 << i
	}

	// keep all that rifraff out
	newVal &= mask >> lowBit

	*cur &= ^mask
	*cur |= newVal << lowBit
}
