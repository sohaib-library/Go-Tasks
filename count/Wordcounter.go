package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	fileData, _ := os.ReadFile("word.txt")

	words := 0
	spaces := 0
	lines := 0
	letters := 0
	special := 0
	paragraphes := 0
	inWord := false

	for _, ch := range string(fileData) {
		letters++

		switch ch {
		case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/', '~', '`':
			special++
		}

		if ch == ' ' {
			spaces++
		}
		if ch == '\n' {
			lines++
		}

		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			inWord = false
		} else if !inWord {
			words++
			inWord = true
		}
		
	}

		fmt.Printf("Total Words %d \n", words)
		fmt.Printf("Total Letter %d \n"  , letters)
		fmt.Printf("Total Lines %d \n"  ,lines)
		fmt.Printf("Total Special %d \n"  ,special)
		fmt.Printf("Total Lines %d \n"  ,lines)
		fmt.Printf("Total Paragraphes %d \n"  ,paragraphes)

	fmt.Println("Execution time:", time.Since(start))
}