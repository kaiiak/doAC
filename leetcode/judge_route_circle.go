func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, v := range moves {
		switch v {
		case rune('R'):
			x++
		case rune('L'):
			x--
		case rune('U'):
			y++
		case rune('D'):
			y--
		}
	}

	if x == 0 && y == 0 {
		return true
	}
	return false
}
