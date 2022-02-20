package flagstr

import (
	"fmt"
	"strings"
)

func New(text string) *flagStr {
	return &flagStr{
		input: text,
	}
}

type flagStr struct {
	input  string
}

func (f *flagStr) Get(arg string) string {
	s := strings.Split(f.input, " ")

	// 分解したsの中からlistに登録した値を探す
	for i := 0; i < len(s) - 1; i++ {
		curText := s[i]
		// if not start '-'
		if !isOptionArg(curText) {
			continue
		}

		if curText == fmt.Sprintf("-%s", arg) {
			return s[i+1]
		}
	}

	return ""
}

func isOptionArg(str string) bool {
	return strings.HasPrefix(str, "-")
}

