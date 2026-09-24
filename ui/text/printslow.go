package text

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

var Stdin = bufio.NewReader(os.Stdin)

func PrintSlow(message string) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return
	}

	skip := make(chan bool, 1)
	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			b, err := Stdin.ReadByte()
			if err != nil {
				return
			}

			if b == ' ' {
				skip <- true
				return
			}
		}
	}()

	for i, lettre := range message {
		select {
		case <-skip:
			fmt.Print(message[i:])
			term.Restore(int(os.Stdin.Fd()), oldState)
			return

		default:
			fmt.Print(string(lettre))
			time.Sleep(50 * time.Millisecond)
		}
	}

	term.Restore(int(os.Stdin.Fd()), oldState)

	select {
	case <-skip:
	case <-done:
	}
}
