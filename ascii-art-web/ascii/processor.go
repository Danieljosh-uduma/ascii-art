package ascii

import (
	"fmt"
	"strings"
)

// TransformText function will process the text using the banner
func TransformText(text, banner string) (string, error) {
	//
	ascii, err := ReadBanner(banner)
	if err != nil {
		return "", fmt.Errorf("Invalid banner name")
	}

	words := strings.Split(text, "\\n")

	// create a string builder for string concatenation
	var builder strings.Builder

	for _, word := range words {
		if word == "" {
			builder.WriteString("\n")
			continue
		}
		for row := 1; row <= 8; row++ {
			for _, char := range word {
				// convert the character to its ASCII value
				// subtract 32 to get the index of the character in the banner
				// add the row number to get the index of the character in the banner
				builder.WriteString(ascii[(int(char)-32)*9+row])
			}
			// add a newline character after each word
			builder.WriteString("\n")
		}
	}

	return builder.String(), nil
}
