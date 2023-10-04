package utils

import (
	"os"
)

func CountLines(file os.File) (int, error) {

	var numberOfLines int
	var currentChar [1]byte

	for {
		n, err := file.Read(currentChar[:])

		// EOF
		if n == 0 {
			break
		}

		if err != nil {
			return 0, err
		}

		if currentChar[0] == '\n' {
			numberOfLines += 1
		}
	}

	return numberOfLines, nil
}
