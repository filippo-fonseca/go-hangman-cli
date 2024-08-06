package main

import (
	"fmt"
	"math/rand"
)

func main() {
	words := []string{"php", "ts", "golang"}
	word := words[rand.Intn(len(words))]
	fmt.Println(word)
}