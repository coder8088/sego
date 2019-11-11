package sego

import (
	"fmt"
	"testing"
)

func TestSegmenter_textToWords(t *testing.T) {
	text := "这是测试ssgs23#4,s;/fq;4.2.?df .2"

	words := textToWords(text)
	for _, w := range words {
		fmt.Printf("%s/", string(w))
	}
}
