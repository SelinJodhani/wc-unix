package utils

import (
	"os"
)

func CountBytes(fileName string) (int, error) {
	file, err := os.Open(fileName)

	defer file.Close()

	if err != nil {
		return 0, err
	}

	fileInfo, err := file.Stat()

	if err != nil {
		return 0, err
	}

	return int(fileInfo.Size()), nil
}
