package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	showLineNum := flag.Bool("n", false, "show line numbers")
	flag.Parse()

	// 残りのファイルパスを取得
	files := flag.Args()

	// 指定なしの場合はエラーを出力
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-n] file1 [file2 ...]\n", os.Args[0])
		os.Exit(1)
	}

	lineNumber := 1

	for _, filename := range files {
		err := processFile(filename, *showLineNum, &lineNumber)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error processing %s: %v\n", filename, err)
			os.Exit(1)
		}
	}
}

func processFile(filename string, showLineNumber bool, lineNumber *int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if showLineNumber {
			fmt.Printf("%6d\t%s\n", *lineNumber, line)
		} else {
			fmt.Println(line)
		}
		(*lineNumber)++
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
