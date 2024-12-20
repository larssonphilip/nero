package terminal

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
	"nero.app/nero/themes"
)

var originalState *term.State

func EnableRawMode() error {
	var err error
	originalState, err = term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	return nil
}

func RestoreTerminal() error {
	return term.Restore(int(os.Stdin.Fd()), originalState)
}

func ReadInput() (byte, error) {
	var buffer [1]byte
	bytesRead, err := os.Stdin.Read(buffer[:])
	if err != nil {
		return 0, err
	}

	if bytesRead == 0 {
		return 0, nil
	}

	return buffer[0], nil
}

func ReadKey() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	b, err := reader.ReadByte()
	if err != nil {
		return "", err
	}

	switch b {
	case 27:
		sequence, err := reader.Peek(2)
		if err != nil {
			return KeyEsc, nil
		}

		if sequence[0] == '[' {
			reader.Discard(2)

			switch sequence[1] {
			case 'A':
				return KeyUp, nil
			case 'B':
				return KeyDown, nil
			case 'C':
				return KeyRight, nil
			case 'D':
				return KeyLeft, nil
			default:
				return "UNKNOWN", nil
			}
		}

		return KeyEsc, nil
	case 127:
		return KeyBackspace, nil
	case 13:
		return KeyEnter, nil
	case 9:
		return KeyTab, nil
	default:
		if b > 31 && b < 127 {
			return string(b), nil
		}
		return "UNKNOWN", nil
	}
}

func ClearScreen() {
	fmt.Print("\x1b[2J\x1b[H")
}

func MoveCursor(row, col int) {
	fmt.Printf("\x1b[%d;%dH", row+1, col+1)
}

func HideCursor() {
	fmt.Print("\x1b[?25l")
}

func ShowCursor() {
	fmt.Print("\x1b[?25h")
}

func ClearLine() {
	fmt.Print("\x1b[2K")
}

func SaveCursorPosition() {
	fmt.Print("\x1b7")
}

func RestoreCursorPosition() {
	fmt.Print("\x1b8")
}

func GetWindowSize() (int, int, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0, err
	}

	return width, height, err
}

func SetTextColor(color int) {
	fmt.Printf("\x1b[%dm", color)
}

func ResetTextColor() {
	fmt.Printf("\x1b[%dm", themes.Default)
}

func EnterFullScreen() {
	fmt.Printf("\x1b[?1049h")
}

func ExitFullScreen() {
	fmt.Printf("\x1b[?1049l")
}

func setBackgroundColor() {}

func resetTextAttributes() {}

func handleResizeSignal() {}

func handleInterruptSignal() {}

func flushInput() {}

func initializeTerminal() {}

func cleanupTerminal() {}

func mapKeys() {}

func isControlKey() {}

func logError(err error) {}
