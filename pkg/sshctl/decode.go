package sshctl

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"unicode"
)

func Decode(r io.Reader) ([]map[string]string, error) {
	sc := bufio.NewScanner(r)

	var list []map[string]string
	m := map[string]string{}
	for sc.Scan() {
		text := removeSpace(sc.Text())

		s := strings.Split(text, " ")

		if len(s) > 2 {
			if s[0] == "ProxyCommand" {
				s[1] = strings.Join(s[1:], " ")
			} else {
				return nil, errors.New("too many args")
			}
		}

		// 空白行で区切り
		if len(s) == 1 {
			// 先頭に空白行があった場合、mに値が入っていないのでスキップ
			if len(m) == 0 {
				continue
			}

			list = append(list, m)
			m = map[string]string{}
			continue
		}

		m[s[0]] = s[1]
	}

	// 最後のmapが残っている場合
	if len(m) > 0 {
		list = append(list, m)
	}

	return list, nil
}

func removeSpace(s string) string {
	for i, r := range s {
		if !unicode.IsSpace(r) {
			return string(s[i:])
		}
	}

	return s
}