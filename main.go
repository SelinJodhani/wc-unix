package main

import (
	"flag"
	"fmt"

	"github.com/SelinJodhani/wc-unix/utils"
)

func main() {
	cflag := flag.Bool("c", false, "Returns number of bytes stored in file if true")
	lflag := flag.Bool("l", false, "Returns number of lines in a file if true")
	wflag := flag.Bool("w", false, "Returns number of words in a file if true")
	mflag := flag.Bool("m", false, "Returns number of characters in a file if true")

	flag.Parse()

	fileName := flag.Args()[0]

	if *cflag {
		byteCount, err := utils.CountBytes(fileName)

		if err != nil {
			panic(err)
		}

		fmt.Println(byteCount, fileName)
		return
	}

	if *lflag {
		lineCount, err := utils.CountLines(fileName)

		if err != nil {
			panic(err)
		}

		fmt.Println(lineCount, fileName)
		return
	}

	if *wflag {
		wordCount, err := utils.CountWords(fileName)

		if err != nil {
			panic(err)
		}

		fmt.Println(wordCount, fileName)
		return
	}

	if *mflag {
		charCount, err := utils.CountChars(fileName)

		if err != nil {
			panic(err)
		}

		fmt.Println(charCount, flag.Args()[0])
		return
	}

	return
}
