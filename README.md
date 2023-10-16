# Challenge 1 - wc unix tool

This challenge corresponds to the first part of the Coding Challenges series by [John Crickett](https://codingchallenges.fyi/challenges/challenge-wc).

## Description

The WC tool is written in `go`. The tool is used to count the number of words, lines, bytes and characters in a file/stdin.

## Build

You can build this project as follows:

```bash
go build ./cmd/ccwc
```

This will make binary file in the root directory named `ccwc`. Binary file of this project is already present inside `./bin/ccwc`.

## Usage

You can use `./bin/ccwc` to run the tool as follows:

```bash
./bin/ccwc [option] filename
```

The following options are supported:

- `-w`: prints the number of words in the file
- `-l`: prints the number of lines in the file
- `-c`: prints the number of bytes in the file
- `-m`: prints the number of characters in the file

The tool can also be used in stdin mode as follows:

```bash
cat filename | ./bin/ccwc [option]
```

## Credits

Thanks to [@jainmohit2001](https://github.com/jainmohit2001) for the thought process and readme.md. Checkout other solutions by him [jainmohit2001/coding-challenges](https://github.com/jainmohit2001/coding-challenges).
