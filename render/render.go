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
	terminal.HideCursor()
	defer terminal.ShowCursor()

	terminalWidth, terminalHeight, err := terminal.GetWindowSize()
	if err != nil {
		fmt.Printf("Error while getting window size: %v\n", err)
	}

	// Calculate the content height with one line reserved for status bar
	e.ContentHeight = terminalHeight - 1
	e.ContentWidth = terminalWidth

	// Adjust the scrolling based on cursor position
	e.Scroll()

	content := e.GetEditorContent()
	lineNumberWidth := calculateLineNumberWidth(content)

	for i := 0; i < e.ContentHeight; i++ {
		fileLineIndex := i + e.RowOffset
		terminal.MoveCursor(i, 0)
		if fileLineIndex < len(content) {
			line := content[fileLineIndex]
			printLine(e.RowOffset, fileLineIndex, lineNumberWidth, line)
		} else {
			fmt.Println("~")
		}
	}

	renderStatusBar(e, terminalWidth)

	cursorX := e.CursorX - lineNumberWidth + 2 - e.ColOffset
	cursorY := e.CursorY - e.RowOffset

	if cursorY >= 0 && cursorY < e.ContentHeight {
		terminal.MoveCursor(cursorY, cursorX)
	}
}

func calculateLineNumberWidth(content []string) int {
	return int(math.Log10(float64(len(content)) + 1))
}

func printLine(rowOffset int, lineNumber, lineNumberWidth int, line string) {
	screenRow := lineNumber - rowOffset
	// Move cursor according to the offset
	terminal.MoveCursor(screenRow, 0)

	terminal.SetTextColor(themes.White)
	fmt.Printf("%*d  ", lineNumberWidth, lineNumber+1)
	terminal.ResetTextColor()

	fmt.Print(line)
	fmt.Print("\r\n")
}

func renderStatusBar(e *editor.Editor, terminalWidth int) {
	fmt.Print("\x1b[7m")

	fileName := e.FilePath
	if fileName == "" {
		fileName = "Untitled"
	}

	modifiedFlag := ""
	if e.Modified {
		modifiedFlag = "[Modified]"
	}

	lineNumber := e.CursorY + 1
	columnNumber := e.CursorX + 1
	totalLines := len(e.GetEditorContent())

	leftStatus := fmt.Sprintf("%s %s - %d lines %s", fileName, modifiedFlag, totalLines, themes.Gray)
	rightStatus := fmt.Sprintf("%d:%d", lineNumber, columnNumber)
	rightStatus = fmt.Sprintf("RowOffset: %d", e.RowOffset)
	padding := terminalWidth - len(leftStatus) - len(rightStatus)
	if padding < 0 {
		padding = 0
	}

	statusBar := fmt.Sprintf("%s%*s%s", leftStatus, padding, " ", rightStatus)

	if len(statusBar) > terminalWidth {
		statusBar = statusBar[:terminalWidth]
	}

	terminal.MoveCursor(e.ContentHeight, 0)

	fmt.Print(statusBar)
	fmt.Print("\x1b[0m")
}
