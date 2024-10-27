package render

import (
	"fmt"

	"nero.app/nero/editor"
	"nero.app/nero/terminal"
)

func RenderScreen(e *editor.Editor) {
	terminal.ClearScreen()
	terminal.MoveCursor(0, 0)

	content := e.GetEditorContent()

	width, _, err := terminal.GetWindowSize()
	if err != nil {
		fmt.Printf("Error while getting window size: %v\n", err)
	}

	for _, line := range content {
		for len(line) > width {
			fmt.Println(line[:width])
			line = line[width:]
		}

		fmt.Println(line)
	}
}
