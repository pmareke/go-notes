package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	file_names := os.Args[1:]

	if len(file_names) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file_name := range file_names {
			file, err := os.Open(file_name)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}

	for line, times := range counts {
		if times > 1 {
			fmt.Printf("Ocurrences: %d\tValue: %s\n", times, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		counts[input.Text()]++
	}
}
