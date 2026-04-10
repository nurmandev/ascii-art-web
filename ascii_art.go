package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var ErrInvalidASCII = errors.New("input contains non-ascii characters")

// GenerateAsciiArt generates ascii art from the given text and banner.
func GenerateAsciiArt(text, banner string) (string, error) {
	if text == "" {
		return "", nil
	}

	file, err := os.Open(banner + ".txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text = strings.ReplaceAll(text, "\r\n", "\n")
	words := strings.Split(text, "\n")

	var result strings.Builder

	allEmpty := true
	for _, word := range words {
		if word != "" {
			allEmpty = false
			break
		}
	}

	// Handle case where text is only newlines
	if allEmpty {
		for i := 0; i < len(words)-1; i++ {
			result.WriteString("\n")
		}
		return result.String(), nil
	}

	for _, word := range words {
		if word == "" {
			result.WriteString("\n")
			continue
		}

		// Print 8 lines for each character
		for i := 0; i < 8; i++ {
			for _, char := range word {
				if char < 32 || char > 126 {
					return "", ErrInvalidASCII
				}

				// The starting index for each character block is 1 + (char - 32) * 9
				// because each character block starts with an empty line (or empty line is considered part of the block above it)
				// Wait, let's verify standard.txt indexing:
				// line 0: ""
				// line 1-8: SPACE
				// line 9: ""
				// line 10-17: !
				// ...
				// (char-32)*9 + 1 -> (0)*9+1 = 1.
				// (33-32)*9 + 1 -> 1*9 + 1 = 10.
				// i goes from 0 to 7. 
				// The logic holds!

				startIndex := int(char-32)*9 + 1
				if startIndex+i < len(lines) {
					result.WriteString(lines[startIndex+i])
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
