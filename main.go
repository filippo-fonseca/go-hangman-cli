package main

import (
	"fmt"
	"strings"

	"github.com/tjarratt/babble"
)

func main() {
	babbler := babble.NewBabbler()
	babbler.Count = 1

	word := string(babbler.Babble())

	lives := 2 * len(word)

	blanks := []string{}

	for range word {
		blanks = append(blanks, "_")
	}

	for {
		fmt.Printf("❤️ %d, Word: %s Letter: ", lives, strings.Join(blanks, " "))

	var input string

	fmt.Scanln(&input)
	input = strings.ToLower(input)

	for _, inputLetter := range input {
		correctGuess := false
		for i, wordLetter := range word {
			if inputLetter == wordLetter {
				blanks[i] = string(inputLetter)
				correctGuess = true
			}
		}

		if !correctGuess {
			 lives--
		}
	}

	if (lives <= 0) {
		fmt.Printf("❤️ 0, Word: %s - sorry, you lost!\n", word)
		break
	}

	if word == strings.Join(blanks, "") {
		fmt.Printf("❤️ %d, Word: %s - you won! Congrats!\n", lives, word)
		break
	}
	}
}