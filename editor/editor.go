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
	terminal.EnterFullScreen()
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

	// Handle key presses
	switch key {
	case terminal.KeyUp:
		if editor.CursorY > 0 {
			editor.CursorY--
		}
	case terminal.KeyDown:
		if editor.CursorY < len(editor.FileContent)-1 {
			editor.CursorY++
		}
	case terminal.KeyRight:
		if editor.CursorX < len(editor.FileContent[editor.CursorY])-1 {
			editor.CursorX++
		}
	case terminal.KeyLeft:
		if editor.CursorX > 0 {
			editor.CursorX--
		}
	case terminal.KeyBackspace:
		line := editor.FileContent[editor.CursorY]
		if editor.CursorX > 0 {
			editor.FileContent[editor.CursorY] = line[:editor.CursorX-1] + line[editor.CursorX:]
			editor.CursorX--
		}
	case terminal.KeyEnter:
		// This might not work as expected
		line := editor.FileContent[editor.CursorY]
		editor.FileContent = append(editor.FileContent[:editor.CursorY+1], editor.FileContent[editor.CursorY:]...)
		editor.FileContent[editor.CursorY] = line[:editor.CursorX]
		editor.FileContent[editor.CursorY+1] = line[editor.CursorX:]
	case terminal.KeyTab:
		line := editor.FileContent[editor.CursorY]
		editor.FileContent[editor.CursorY] = line[:editor.CursorX] + "    " + line[editor.CursorX:]
		editor.CursorX += 4
	case terminal.KeyEsc:
		terminal.ExitFullScreen()
		os.Exit(0)
	default:
		line := editor.FileContent[editor.CursorY]
		editor.FileContent[editor.CursorY] = line[:editor.CursorX] + string(key) + line[editor.CursorX:]
		editor.CursorX++
	}

	return nil
}
