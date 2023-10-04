package utils

import "os"

func CountBytes(file os.File) (int, error) {
	var byteCountInFile int

	buffer := make([]byte, 4096)

	for {
		n, err := file.Read(buffer)

		// EOF
		if n == 0 {
			break
		}

		if err != nil {
			return 0, err
		}

		byteCountInFile += n
	}

	return byteCountInFile, nil
}
