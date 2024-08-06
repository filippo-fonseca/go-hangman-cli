package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	words := []string{"php", "ts", "golang"}
	word := words[rand.Intn(len(words))]

	blanks := []string{}

	for range word {
		blanks = append(blanks, "_")
	}

	fmt.Printf("Word: %s Letter: ", strings.Join(blanks, " "))

	var input string

	fmt.Scanln(&input)
	fmt.Println(input)

	for _, inputLetter := range input {
		for i, wordLetter := range word {
			if inputLetter == wordLetter {
				blanks[i] = string(inputLetter)
			}
		}
	}

	fmt.Println(blanks)
}