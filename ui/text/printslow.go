package text

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

func PrintSlow(message string) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	skip := make(chan bool)

	go func() {
		var key [1]byte

		for {
			os.Stdin.Read(key[:])

			if key[0] == ' ' {
				skip <- true
				return
			}
		}
	}()

	for i, lettre := range message {
		select {
		case <-skip:
			fmt.Print(message[i:])
			return

		default:
			fmt.Print(string(lettre))
			time.Sleep(50 * time.Millisecond)
		}
	}
}
