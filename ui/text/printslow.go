package text

import (
	"fmt"
	"time"
)

func PrintSlow(message string) {
	for _, lettre := range message {
		fmt.Print(string(lettre))
		time.Sleep(30 * time.Millisecond)
	}
}
