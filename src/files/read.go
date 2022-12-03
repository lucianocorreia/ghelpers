// Copyright 2022 Luciano Correia.  All rights reserved.

package files

import (
	"bufio"
	"os"
)

// ReadLines returns a list of strings from a given file path
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines, sc.Err()
}
