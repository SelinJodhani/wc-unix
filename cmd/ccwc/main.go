package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/SelinJodhani/wc-unix/internal/count"
)

func main() {
	var bytesFlag, linesFlag, wordsFlag, charsFlag bool

	flag.BoolVar(&bytesFlag, "c", false, "Returns number of bytes stored in file if true")
	flag.BoolVar(&bytesFlag, "bytes", false, "Returns number of bytes stored in file if true")

	flag.BoolVar(&linesFlag, "l", false, "Returns number of lines in a file if true")
	flag.BoolVar(&linesFlag, "lines", false, "Returns number of lines in a file if true")

	flag.BoolVar(&wordsFlag, "w", false, "Returns number of words in a file if true")
	flag.BoolVar(&wordsFlag, "words", false, "Returns number of words in a file if true")

	flag.BoolVar(&charsFlag, "m", false, "Returns number of characters in a file if true")
	flag.BoolVar(&charsFlag, "chars", false, "Returns number of characters in a file if true")

	flag.Parse()

	var r io.Reader
	var fileName string
	var cleanup func()

	if len(flag.Args()) == 0 {
		r = os.Stdin
	} else {
		fileName = flag.Args()[0]
		file, err := os.Open(fileName)

		if os.IsNotExist(err) {
			fmt.Println(err)
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
	}

	if r == os.Stdin {
		tempFile, err := os.CreateTemp("", "stdin-tempfile-")

		if err != nil {
			fmt.Println("Error creating temporary file:", err)
			os.Exit(1)
		}

		defer tempFile.Close()
		defer os.Remove(tempFile.Name())

		_, err = io.Copy(tempFile, os.Stdin)

		if err != nil {
			fmt.Println("Error copying stdin to temporary file:", err)
			os.Exit(1)
		}

		fileName = tempFile.Name()
		cleanup = func() {
			os.Remove(tempFile.Name()) // Ensure temporary file is deleted
		}
	}

	var wg sync.WaitGroup

	byteResult := make(chan int, 1)
	wordResult := make(chan int, 1)
	lineResult := make(chan int, 1)
	charResult := make(chan int, 1)

	if !bytesFlag && !linesFlag && !wordsFlag && !charsFlag {
		bytesFlag = true
		linesFlag = true
		wordsFlag = true
	}

	if bytesFlag {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(byteResult)

			bytes, err := count.Bytes(fileName)

			if err != nil {
				panic(err)
			}

			byteResult <- bytes
		}()
	}

	if linesFlag {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(lineResult)

			lines, err := count.Lines(fileName)

			if err != nil {
				panic(err)
			}

			lineResult <- lines
		}()
	}

	if wordsFlag {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(wordResult)

			words, err := count.Words(fileName)

			if err != nil {
				panic(err)
			}

			wordResult <- words
		}()
	}

	if charsFlag {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(charResult)

			chars, err := count.Chars(fileName)

			if err != nil {
				panic(err)
			}

			charResult <- chars
		}()
	}

	wg.Wait()

	fmt.Print(" ")

	if linesFlag {
		fmt.Print(" ", <-lineResult)
	}

	if wordsFlag {
		fmt.Print(" ", <-wordResult)
	}

	if bytesFlag {
		fmt.Print(" ", <-byteResult)
	}

	if charsFlag {
		fmt.Print(" ", <-charResult)
	}

	if r != os.Stdin {
		fmt.Println(" ", fileName)
	} else {
		fmt.Println(" ")
	}

	if cleanup != nil {
		cleanup()
	}
}
