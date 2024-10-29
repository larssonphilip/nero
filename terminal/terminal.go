package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
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
	b := make([]byte, 1)
	_, err := os.Stdin.Read(b)
	if err != nil {
		return "", err
	}

	// Check if the input is an escape sequence
	if b[0] == 27 {
		sequence := make([]byte, 2)
		os.Stdin.Read(sequence)

		// Double press ESC to exit
		if len(sequence) > 1 && sequence[1] == 0 {
			return KeyEsc, nil
		}

		// Interpret the escape sequence as an arrow key
		switch string(sequence) {
		case "[A":
			return KeyUp, nil
		case "[B":
			return KeyDown, nil
		case "[C":
			return KeyRight, nil
		case "[D":
			return KeyLeft, nil
		default:
			return "UNKNOWN", nil
		}
	}

	// Return the single character as a string
	return string(b), nil
}

func ClearScreen() {
	fmt.Print("\x1b[2J\x1b[H")
}

func MoveCursor(row, col int) {
	fmt.Printf("\x1b[%d;%dH", col+1, row+1)
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

func setTextColor() {}

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
