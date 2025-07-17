package main

import (
	"fmt"
	"os"
)

func listFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("ディレクトリの読み込みに失敗しました: %w", err)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}

func main() {
	dir := "./Problems"
	files, err := listFiles(dir)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
