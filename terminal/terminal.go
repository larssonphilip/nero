package terminal

import (
	"os"

	"golang.org/x/term"
)

func enableRawMode() (*term.State, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	return oldState, nil
}

func restoreTerminal(oldState *term.State) error {
	return term.Restore(int(os.Stdin.Fd()), oldState)
}

func readInput() (byte, error) {
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

func readKey() (string, error) {
	b, err := readInput()
	if err != nil {
		return "", err
	}

	// Escape sequence
	if b == 27 {
		sequence := make([]byte, 2)
		bytesRead, err := os.Stdin.Read(sequence)
		if bytesRead == 0 || err != nil {
			return "ESC", err
		}

		switch string(sequence) {
		case "[A":
			return "UP", nil
		case "[B":
			return "DOWN", nil
		case "[C":
			return "RIGHT", nil
		case "[D":
			return "LEFT", nil
		default:
			return "UNKNOWN", nil
		}
	}

	return string(b), nil
}
