package editor

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"nero.app/nero/terminal"
)

func HandleResize(editor *Editor) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGWINCH)
	go func() {
		for range c {
			_, _, err := terminal.GetWindowSize()
			if err != nil {
				fmt.Printf("Error while getting window size: %v\n", err)
			}

			// render.RenderScreen(editor)
		}
	}()
}
