package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [directory path]")
		os.Exit(1)
	}

	dirPath := os.Args[1]
	fmt.Println("Organizing files in:", dirPath)
}

func readFiles(dirPath string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	return files, nil

}

func getCategory(fileName string) string {
	ext := filepath.Ext(fileName)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return "Images"
	case ".mp4", ".mkv", "mov", ".avi":
		return "Videos"
	case ".mp3":
		return "Music"
	case ".pdf", ".doc", ".docx", ".txt", ".ppt", ".pptx", ".xls", ".xlsx":
		return "Documents"
	default:
		return "Others"
	}
}
