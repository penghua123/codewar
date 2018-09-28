package main

import (
	"fmt"
)

func main() {
	keyPressCount := 2017
	key := WhichNote(keyPressCount)
	fmt.Println(key)
	key = WhichNoteBetter(keyPressCount)
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

func WhichNoteBetter(keyPressCount int) string {
	keyType := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	key := keyType[(keyPressCount-1)%88%len(keyType)]
	return key
}
