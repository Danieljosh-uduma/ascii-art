package main

import "testing"

func TestReadCharacter(t *testing.T) {
	_, err := readCharacters()

	if err != nil {
		t.Error("Function was not able to read characters")
	}
}
