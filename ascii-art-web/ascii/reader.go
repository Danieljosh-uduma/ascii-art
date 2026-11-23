package ascii

import (
	"bufio"
	"os"
)

// ReadBanner opens a text file, reads it line by line, and
// returns a slice of strings that contains the file content
func ReadBanner(txtfile string) ([]string, error) {
	// This opens the file for reading except there is an error
	file, errorOccured := os.Open(txtfile)
	if errorOccured != nil {
		return nil, errorOccured
	}
	defer file.Close()

	var lines []string

	// scanner is initialized to read the file line by line
	scanner := bufio.NewScanner(file)

	// Moves line by line to the EOF while adding each line's text to the 'lines' slice
	for scanner.Scan() {
		// Using the append function elements of the text are added to the list of strings
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
