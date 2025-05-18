package main

import (
	"fmt"
	"strings"
)

func main() {
	var words []string
	for range 10 {
		words = append(words, randomWord())
	}
	result := strings.Join(words, "-")
	fmt.Println(result)
}
