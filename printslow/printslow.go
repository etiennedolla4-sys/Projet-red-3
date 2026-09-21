package printslow

import (
	"fmt"
	"time"
)

func PrintSlow(s string) {
	for _, char := range s {
		fmt.Printf("%c", char)
		time.Sleep(50 * time.Millisecond)
	}
}
