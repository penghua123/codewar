package pianoKata

func BlackOrWhiteKey(keyPressCount int) string {
	color := make(map[int]string)
	color[1] = "white"
	color[2] = "black"
	color[3] = "white"
	color[4] = "black"
	color[5] = "white"
	color[6] = "white"
	color[7] = "black"
	color[8] = "white"
	color[9] = "black"
	color[10] = "white"
	color[11] = "black"
	color[0] = "white"
	keyLoopCount := keyPressCount % 88
	if keyLoopCount <= 3 {
		return color[keyLoopCount]
	}
	if keyLoopCount > 3 && keyLoopCount < 88 {
		pos := (keyLoopCount - 3) % 12
		return color[pos]
	}
	return "Error!!! There is missing data"
}
