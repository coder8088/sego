package sego

import (
	"github.com/tedux/sego/dict"
	"unicode"
	"unicode/utf8"
)

type Segmenter interface {
	SetDict(dict.Dict)
	Segment(string)
}

type segmenter struct {
	dict dict.Dict
}

func New() *segmenter {
	// 加载一个默认字典
	return &segmenter{dict: dict.NewDict()}
}

func (seg *segmenter) SetDict(dict dict.Dict) {
	seg.dict = dict
}

func (seg *segmenter) Segment(text string) {
	// wordBytesArray := textToWords(text)

}

func textToWords(text string) [][]byte {
	byteArray := []byte(text)
	output := make([][]byte, 0, len(byteArray)/3)
	current := 0
	inAlphanumeric := true
	alphanumericBegin := 0
	for current < len(byteArray) {
		r, size := utf8.DecodeRune(byteArray[current:])
		if size <= 2 && (unicode.IsLetter(r) || unicode.IsNumber(r) || isConnectorRune(r)) {
			// 当前是拉丁字母或者数字
			if !inAlphanumeric {
				alphanumericBegin = current
				inAlphanumeric = true
			}
		} else {
			if inAlphanumeric {
				inAlphanumeric = false
				if current != 0 {
					output = append(output, toLower(byteArray[alphanumericBegin:current]))
				}
			}
			//if !unicode.IsPunct(r) && !unicode.IsSpace(r) {
			output = append(output, byteArray[current:current+size])
			//}
		}
		current += size
	}
	// 处理最后一个字元是英文的情况
	if inAlphanumeric {
		if current != 0 {
			output = append(output, toLower(byteArray[alphanumericBegin:current]))
		}
	}

	return output
}

func isConnectorRune(r rune) bool {
	if r == '#' || r == '&' || r == '+' || r == '-' || r == '.' || r == '@' || r == '_' || r == ',' || r == '*' {
		return true
	}
	return false
}

func toLower(text []byte) []byte {
	output := make([]byte, len(text))
	for i, b := range text {
		if b >= 'A' && b <= 'Z' {
			output[i] = b - 'A' + 'a'
		} else {
			output[i] = b
		}
	}
	return output
}
