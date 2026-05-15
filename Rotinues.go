package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"time"
)

type result struct{
	word int
	lineno int
	spaces int 
	vawles int 
	special int
	letters int 
}



func main() {

	time := time.Now()

	Filedata, _:= os.ReadFile("word.txt")
	content := string(Filedata)

	TotalLen := len(content)

	
	worker := 4
   chunkSize := (TotalLen + worker - 1) / worker

   returnCh := make(chan result, worker)

   // chunk data 

   for i := 0; i < worker; i++ {

	start := i*chunkSize
	end := chunkSize + start

	if i == worker -1 {

       end = TotalLen
		
	}
	wg.Add(1)

	go func(chunk) {

		defer wg.Done()

		returnCh <- result{}
		inWord := false

		for _, b := range chunk {
				ch := int(b)
				

				switch ch {
				case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/', '~', '`':
					res.special++
				}

				if ch == ' ' {
					res.spaces++
				}
				if ch == '\n' {
					res.lines++
				}

				if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
					inWord = false
				} else if !inWord {
					res.words++
					inWord = true
				}
			}

			resultChan <- res
		}(fileData[start:end])

		
    }
}
   
    go func() {
    	wg.Wait()
    	close(resultCh)
    
    }()

	finalWords := 0
	finalSpaces := 0
	finalLines := 0
	finalLetters := 0
	finalSpecial := 0

	for res := range resultChan {
		finalWords += res.words
		finalSpaces += res.spaces
		finalLines += res.lines
		finalLetters += res.letters
		finalSpecial += res.special
	}

	fmt.Printf("Total Words %d \n", finalWords)
	fmt.Printf("Total Letter %d \n", finalLetters)
	fmt.Printf("Total Lines %d \n", finalLines)
	fmt.Printf("Total Special %d \n", finalSpecial)

	fmt.Println("Execution time:", time.Since(start))

}
