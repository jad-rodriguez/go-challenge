package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the text file
	file, err := os.Open("properties.txt")
	// check for any error opening file
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	// Used bufio to create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	// Split each line
	scanner.Split(bufio.ScanLines)
	// Assign a variable to contain lines as strings
	var txtlines []string
	// Use Scan
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	// file.Close()

	// Loop through each line
	// Use RemoveDuplicate function
	// Print eachline
	for _, eachline := range txtlines {
		RemoveDuplicates(&txtlines)
		fmt.Println(eachline)
	}
}

// Function to remove duplicates
func RemoveDuplicates(lines *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *lines {
		if !found[x] {
			found[x] = true
			(*lines)[j] = (*lines)[i]
			j++
		}
	}
	*lines = (*lines)[:j]
}
