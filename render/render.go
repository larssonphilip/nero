package render

import (
	"fmt"
	"math"

	"nero.app/nero/editor"
	"nero.app/nero/terminal"
	"nero.app/nero/themes"
)

func RenderScreen(e *editor.Editor) {
	terminal.ClearScreen()
	terminal.MoveCursor(0, 0)

	content := e.GetEditorContent()

	width, _, err := terminal.GetWindowSize()
	if err != nil {
		fmt.Printf("Error while getting window size: %v\n", err)
	}

	lineNumberWidth := calculateLineNumberWidth(content)

	for lineNumber, line := range content {
		for len(line) > width {
			printLine(lineNumber, lineNumberWidth, line[:width])
			line = line[width:]
		}

		printLine(lineNumber, lineNumberWidth, line)
	}

	terminal.MoveCursor(e.CursorX+lineNumberWidth+2, e.CursorY)
}

func calculateLineNumberWidth(content []string) int {
	return int(math.Log10(float64(len(content))) + 1)
}

func printLine(lineNumber, lineNumberWidth int, line string) {
	terminal.SetTextColor(themes.Red)
	fmt.Printf("%*d  ", lineNumberWidth, lineNumber+1)
	terminal.ResetTextColor()
	fmt.Print(line + "\r\n")
}
