package base

func Max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func SwapRune(data []rune, x, y int) {
	temp := data[x]
	data[x] = data[y]
	data[y] = temp
}
