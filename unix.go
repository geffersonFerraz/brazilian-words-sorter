//go:build !windows
// +build !windows

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

	}

	return result

}
