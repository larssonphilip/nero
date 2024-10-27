package editor

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
