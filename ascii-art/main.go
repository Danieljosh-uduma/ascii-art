package main

import "fmt"

type asciiArt [][]string

// main function get the word(s) from the os.Args
// read the characters, then print the equivalent ascii character.
func main() {
	word, err := getWord()
	if err != nil {
		fmt.Println(err)
		return
	}

	characters, err := readCharacters()
	if err != nil {
		fmt.Println(err)
		return
	}

	printWord(&characters, word)
}
