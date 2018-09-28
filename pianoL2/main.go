package main

import (
	"fmt"
)

func main() {
	key := WhichNote(101)
	fmt.Println(key)
}

func WhichNote(keyPressCount int) string {
	keyRealCount := keyPressCount % 88
	if keyRealCount == 0 {
		keyRealCount += 4
	}
	keyType := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	key := keyType[(keyRealCount-1)%len(keyType)]
	return key
}
