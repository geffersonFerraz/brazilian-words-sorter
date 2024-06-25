package bws

import (
	"strings"
	"testing"
)

func TestSortWords(t *testing.T) {
	bws := BrazilianWords(3, ",")
	result := bws.Sort()
	if len(strings.Split(result, ",")) != 3 {
		t.Errorf("Sort() = %v, expected %v", len(result), 3)
	}
}
