package main

import (
	"fmt"
	"os"
	"strings"
)

// getWord function checks if os arguments length is 2
// return the first argument
func getWord() (string, error) {
	if len(os.Args) != 2 {
		return "", fmt.Errorf(`Enter a string you want to display. e.g "Hello World?!"`)
	}
	return os.Args[1], nil
}

// readCharacters function from the ascii standard file
// returns err if not found
// return [][]strings containing each character from whitespace to ~
func readCharacters() (asciiArt, error) {
	bytes, err := os.ReadFile("templates/standard.txt")

	if err != nil {
		return nil, fmt.Errorf("File not found.")
	}

	var characters asciiArt
	lines := strings.Split(string(bytes), "\n")

	for i := 1; i < len(lines); i += 8 {
		end := i + 8
		if end > len(lines) {
			end = len(lines)
		}
		characters = append(characters, lines[i:end])
		i++ // skip the newline
	}

	return characters, nil
}

// printWord functions print ascii words
// prints a new line when \n character is found
func printWord(char *asciiArt, words string) {
	if words == "" { // handles empty string
		return
	}
	//? I dont know
	word := strings.Split(words, "\n")

	for _, r := range word {
		if r == "" {
			fmt.Println() // handles the "\n" character
			continue
		}
		print(char, r)
	}
}

// print function print finds and prints ascii characters
func print(char *asciiArt, word string) {
	characters := *char

	for row := 0; row < 8; row++ {
		for _, ch := range word {

			index := ch - 32
			if index >= 0 && index <= rune(len(characters)) {
				fmt.Print(characters[index][row])
			}
		}
		fmt.Println()
	}
}
