package editor

import (
	"os"

	"nero.app/nero/terminal"
)

type Editor struct {
	FileContent []string
	CursorX     int
	CursorY     int
}

func InitializeEditor(fileContent []string) *Editor {
	return &Editor{
		FileContent: fileContent,
		CursorX:     0,
		CursorY:     0,
	}
}

func (editor *Editor) GetEditorContent() []string {
	return editor.FileContent
}

func (editor *Editor) ProcessKeyPress() error {
	key, err := terminal.ReadKey()
	if err != nil {
		return err
	}

	// Handle key-based cursor movement
	switch key {
	case "UP":
		if editor.CursorY > 0 {
			editor.CursorY--
		}
	case "DOWN":
		if editor.CursorY < len(editor.FileContent)-1 {
			editor.CursorY++
		}
	case "RIGHT":
		if editor.CursorX < len(editor.FileContent[editor.CursorY])-1 {
			editor.CursorX++
		}
	case "LEFT":
		if editor.CursorX > 0 {
			editor.CursorX--
		}
	case "ESC":
		os.Exit(0)
	}

	return nil
}
