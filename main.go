package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/SelinJodhani/wc-unix/utils"
)

func main() {
	cflag := flag.Bool("c", false, "Returns number of bytes stored in file if true")
	lflag := flag.Bool("l", false, "Returns number of lines in a file if true")
	wflag := flag.Bool("w", false, "Returns number of words in a file if true")
	mflag := flag.Bool("m", false, "Returns number of characters in a file if true")

	flag.Parse()

	fileName := flag.Args()[0]

	var wg sync.WaitGroup

	byteResult := make(chan int, 1)
	wordResult := make(chan int, 1)
	lineResult := make(chan int, 1)
	charResult := make(chan int, 1)

	if !*cflag && !*lflag && !*wflag && !*mflag {
		*cflag = true
		*lflag = true
		*wflag = true
	}

	if *cflag {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(byteResult)

			bytes, err := utils.CountBytes(fileName)

			if err != nil {
				panic(err)
			}

			byteResult <- bytes
		}()
	}

	if *lflag {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(lineResult)

			lines, err := utils.CountLines(fileName)

			if err != nil {
				panic(err)
			}

			lineResult <- lines
		}()
	}

	if *wflag {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(wordResult)

			words, err := utils.CountWords(fileName)

			if err != nil {
				panic(err)
			}

			wordResult <- words
		}()
	}

	if *mflag {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(charResult)

			chars, err := utils.CountChars(fileName)

			if err != nil {
				panic(err)
			}

			charResult <- chars
		}()
	}

	wg.Wait()

	fmt.Print(" ")

	if *lflag {
		fmt.Print(<-lineResult)
	}

	if *wflag {
		fmt.Print(<-wordResult)
	}

	if *cflag {
		fmt.Print(<-byteResult)
	}

	if *mflag {
		fmt.Print(<-charResult)
	}

	fmt.Println(" ", fileName)
	return
}
