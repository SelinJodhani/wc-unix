package utils

import (
	"bufio"
	"io"
	"os"
)

func CountChars(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return 0, err
	}

	count := 0
	reader := bufio.NewReader(file)

	for {
		_, _, err := reader.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, err
		}

		count += 1
	}

	return count, nil
}
