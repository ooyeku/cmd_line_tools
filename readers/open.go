package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		_, err := fmt.Fprintln(os.Stderr, "Error: Incorrect number of arguments.")
		if err != nil {
			return
		}
		_, err = fmt.Fprintf(os.Stderr, "Usage: %s <file_path>\n", os.Args[0])
		if err != nil {
			return
		}
		os.Exit(1)
	}

	filePath := os.Args[1]
	ext := filepath.Ext(filePath)

	var cmd *exec.Cmd
	switch ext {
	case ".txt":
		cmd = exec.Command("open", filePath)
	case ".pdf":
		cmd = exec.Command("open", filePath)
	case ".csv":
		cmd = exec.Command("open", "-a", "Terminal", filePath)
	default:
		_, err := fmt.Fprintln(os.Stderr, "Error: Unsupported file type.")
		if err != nil {
			return
		}
		os.Exit(1)
	}

	err := cmd.Run()
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, "Error:", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
