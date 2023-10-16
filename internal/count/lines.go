package count

import (
	"bufio"
	"os"
)

func Lines(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return 0, err
	}

	count := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		count += 1
	}

	return count, nil
}
