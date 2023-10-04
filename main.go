package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SelinJodhani/wc-unix/utils"
)

func main() {

	cflag := flag.Bool("c", false, "Returns number of bytes stored in file if true")
	lflag := flag.Bool("l", false, "Returns number of lines in a file if true")

	flag.Parse()

	file, err := os.Open(flag.Args()[0])

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	if *cflag {

		byteCount, err := utils.CountBytes(*file)

		if err != nil {
			panic(err)
		}

		fmt.Println(byteCount, flag.Args()[0])

		return
	}

	if *lflag {
		lineCount, err := utils.CountLines(*file)

		if err != nil {
			panic(err)
		}

		fmt.Println(lineCount, flag.Args()[0])

		return
	}
}
