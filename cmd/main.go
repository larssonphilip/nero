package main

import (
	"fmt"
	"os"

	"nero.app/nero/editor"
	"nero.app/nero/file"
	"nero.app/nero/render"
	"nero.app/nero/terminal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: nero <filename>")
		os.Exit(1)
	}

	if err := terminal.EnableRawMode(); err != nil {
		fmt.Printf("Error enabling raw mode: %v\n", err)
		os.Exit(1)
	}

	defer terminal.RestoreTerminal()

	filePath := os.Args[1]

	fileContent, err := file.LoadFile(filePath)
	if err != nil {
		fmt.Printf("Error loading file: %v\n", err)
		os.Exit(1)
	}

	e := editor.InitializeEditor(filePath, fileContent)

	for {
		render.RenderScreen(e)
		if err := e.ProcessKeyPress(); err != nil {
			fmt.Printf("Error processing key press: %v\n", err)
			break
		}
	}
}
