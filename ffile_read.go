// Copyright 2022 Luciano Correia.  All rights reserved.

package ghhelpers

import (
	"bufio"
	"io"
	"os"
)

// FileReadLines returns a list of strings from a given file path
func FileReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return readlinesFromReader(file)
}

// readlinesFromReader reads all lines from a given reader
func readlinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines, sc.Err()
}
