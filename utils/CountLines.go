package utils

import (
	"bufio"
	"os"
)

func CountLines(fileName string) (int, error) {
	count := 0
	file, err := os.Open(fileName)

	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		count += 1
	}

	return count, nil
}
