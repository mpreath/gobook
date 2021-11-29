package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, fileNames := range counts {
		total_count := 0
		occurences := ""
		sep := ""
		for fileName, count := range fileNames {
			total_count += count
			occurences += sep + fileName + "[" + strconv.Itoa(count) + "]"
			sep = " "
		}
		fmt.Printf("%d\t%s\t(%s)\n", total_count, line, occurences)
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][f.Name()]++
	}
	err := input.Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	}
}
