package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the file for reading
	file, err := os.Open("Important people.md")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a slice to hold the modified lines
	var lines []string

	// Read each line and modify if it starts with "#"
	scanner := bufio.NewScanner(file)
	for _, scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			line = "#" + line
		}
		lines = append(lines, line)
	}

	// Check for any errors encountered while reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Open the file for writing (truncate the file first)
	file, err = os.OpenFile("Important people.md", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file for writing:", err)
		return
	}
	defer file.Close()

	// Write the modified lines back to the file
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// Ensure all data is flushed to the file
	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing data to file:", err)
	}
}
