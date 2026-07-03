package analyzer_service

import (
	"fmt"
	"sync"
	"time"
	"web-server/models"

	// "go.uber.org/mock/mockgen/model"
)

type result models.WordCountResponse ;

func Filecontext(Id int, content string) (int, int, int, int, int) {

	totalLen := len(content)

	worker := Id
	chunkSize := (totalLen + worker - 1) / worker

	resultChan := make(chan models.WordCountResponse, worker)

	var wg sync.WaitGroup

	startTime := time.Now()

	for i := range worker {
		start := i * chunkSize
		end := start + chunkSize

		if i == worker-1 {
			end = totalLen
		}

		wg.Add(1)

		go func(chunk string) {
			defer wg.Done()

			var res models.WordCountResponse
			inWord := false

			for j := 0; j < len(chunk); j++ {
				ch := chunk[j]

				switch ch {
				case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=',
					'{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/', '~', '`':
					res.SpecialCharCount++
				}

				if ch == ' ' {
					res.SpaceCount++
				}
				if ch == '\n' {
					res.LineCount++
				}

				if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
					inWord = false
				} else {
					if !inWord {
						res.WordCount++
						inWord = true
					}
					if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
						res.LetterCount++
					}
				}
			}

			resultChan <- res
		}(content[start:end])
	}

	wg.Wait()
	close(resultChan)

	finalWords := 0
	finalSpaces := 0
	finalLines := 0
	finalLetters := 0
	finalSpecial := 0

	for res := range resultChan {
		finalWords += res.WordCount
		finalSpaces += res.SpaceCount
		finalLines += res.LineCount
		finalLetters += res.LetterCount
		finalSpecial += res.SpecialCharCount
	}

	elapsed := time.Since(startTime)

	fmt.Printf("Total Words:   %d\n", finalWords)
	fmt.Printf("Total Letters: %d\n", finalLetters)
	fmt.Printf("Total Lines:   %d\n", finalLines)
	fmt.Printf("Total Spaces:  %d\n", finalSpaces)
	fmt.Printf("Total Special: %d\n", finalSpecial)
	fmt.Printf("Total Special: %d\n", finalSpecial)

	fmt.Println("Execution time:", elapsed)

	return finalLetters, finalWords, finalLines, finalSpaces, finalSpecial

}
