package main

import (
	"fmt"
	"os"
	"time"
)

func countWords(content string) int {

	// Wrod Counting

	count := 0
	inWord := false

	for _, ch := range content {
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			inWord = false
		} else {
			if !inWord {
				count++
				inWord = true
			}
		}
	}
	return count

}

// space counting

func spaceCounting(constant string) int {

	spaceCount := 0

	for _, ch := range constant {

		if ch == ' ' {
			spaceCount++
		}

	}
	return spaceCount

}

// Line Counting
func Linecounter(content string) int {

	Lineno := 0

	for _, ch := range content {

		if ch == '\n' {

			Lineno++
		}

	}
	return Lineno
}

// Letters
func Letterscounter(constant string) int {
	return len(constant)
}

//No Of  Special Characters

func Spaceialchars(constant string) int {

	Special := 0

	for _, ch := range constant {

		switch ch {
		case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/', '~', '`':
			Special++
		}

	}
	return Special

}

func main() {

	start := time.Now()

	data, _ := os.ReadFile("word.txt")

	content := string(data)

	fmt.Println("Total Words  :", countWords(content))
	fmt.Println("Total Spaces  :", spaceCounting(content))
	fmt.Println("Total Lines  :", Linecounter(content))
	fmt.Println("Total Letters :", Letterscounter(content))
	fmt.Println("Total Special Characters :", Spaceialchars(content))

	fmt.Println("Execution of Time:", time.Since(start))
}