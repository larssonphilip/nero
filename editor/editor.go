package editor

import (
	"os"

	"nero.app/nero/terminal"
)

type Editor struct {
	FilePath      string
	FileContent   []string
	CursorX       int
	CursorY       int
	Modified      bool
	RowOffset     int
	ColOffset     int
	ContentWidth  int
	ContentHeight int
}

func InitializeEditor(filePath string, fileContent []string) *Editor {
	terminal.EnterFullScreen()
	terminalWidth, terminalHeight, err := terminal.GetWindowSize()
	if err != nil {
		terminal.ExitFullScreen()
		os.Exit(0)
	}
	return &Editor{
		FilePath:      filePath,
		FileContent:   fileContent,
		CursorX:       0,
		CursorY:       0,
		Modified:      false,
		RowOffset:     0,
		ColOffset:     0,
		ContentWidth:  terminalWidth,
		ContentHeight: terminalHeight - 1,
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
			editor.Modified = true
		}
	case terminal.KeyEnter:
		// This might not work as expected
		line := editor.FileContent[editor.CursorY]
		editor.FileContent = append(editor.FileContent[:editor.CursorY+1], editor.FileContent[editor.CursorY:]...)
		editor.FileContent[editor.CursorY] = line[:editor.CursorX]
		editor.FileContent[editor.CursorY+1] = line[editor.CursorX:]
		editor.Modified = true
	case terminal.KeyTab:
		line := editor.FileContent[editor.CursorY]
		editor.FileContent[editor.CursorY] = line[:editor.CursorX] + "    " + line[editor.CursorX:]
		editor.CursorX += 4
		editor.Modified = true
	case terminal.KeyEsc:
		terminal.ExitFullScreen()
		os.Exit(0)
	default:
		line := editor.FileContent[editor.CursorY]
		editor.FileContent[editor.CursorY] = line[:editor.CursorX] + string(key) + line[editor.CursorX:]
		editor.CursorX++
		editor.Modified = true
	}

	return nil
}

func (editor *Editor) Scroll() {
	if editor.CursorY < editor.RowOffset {
		editor.RowOffset = editor.CursorY
	}

	if editor.CursorY >= editor.RowOffset+editor.ContentHeight {
		editor.RowOffset = editor.CursorY - editor.ContentHeight + 1
	}
}

func saveFileAs(filePath string, content []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, line := range content {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func (editor *Editor) saveCurrentFile() error {
	return saveFileAs(editor.FilePath, editor.FileContent)
}
