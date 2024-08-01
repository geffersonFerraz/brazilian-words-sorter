//go:build windows
// +build windows

package bws

import (
	"math/rand"
	"strings"
	"time"
)

func (b *brazilianWords) Sort() string {
	last := len(wordList) - 1
	result := ""
	for i := 1; i <= b.totalWords; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		rng := rand.New(source)
		result = result + strings.ToLower(wordList[rng.Intn(last)])
		if i < b.totalWords {
			result = result + b.separator
		}
		// REASON: Timer Resolution
		time.Sleep(20 * time.Millisecond)
	}

	return result

}

// Timer Resolution - https://pkg.go.dev/time
/*
Timer resolution varies depending on the Go runtime, the operating system and the underlying hardware.
On Unix, the resolution is approximately 1ms. On Windows, the default resolution is approximately 16ms,
but a higher resolution may be requested using golang.org/x/sys/windows.TimeBeginPeriod.
*/
