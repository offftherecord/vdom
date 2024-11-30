package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Define a regular expression for a valid domain
	domainRegex := regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)

	// Create a scanner to read lines from stdin
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line matches the domain format
		if domainRegex.MatchString(line) {
			fmt.Println(line)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
