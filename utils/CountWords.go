package utils

import (
	"bufio"
	"os"
	"strings"
)

func CountWords(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return 0, err
	}

	count := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		count += len(words)
	}

	return count, nil
}
