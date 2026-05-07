package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordcounter() {

}

func main() {

	file, _ := os.Open("counter.txt")
    defer file.Close()

   r := bufio.NewReader(file)

for {

  file, _, err := r.ReadLine()
  
  if len(file) > 0 {

    fmt.Printf("%q\n", file)

  }

  if err != nil {

	fmt.Printf("this is an error %v\n" ,err)

    break

  }

 }

}
