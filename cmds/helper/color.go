package helper

import (
	"fmt"
	"io"
	"unicode"
)

const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textPurple
	textCyan
	textWhite
)

// 这里提供终端的ASCII输出
type colorWriter struct {
	w io.Writer
}

// Write 实现了 io.Writer 的 Write 接口
func (cw *colorWriter) Write(p []byte) (int, error) {
	return cw.w.Write(p)
}

func bold(message string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[21m", message)
}

func red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}

func capitalize(message string) string {
	var upStr string
	vm := []rune(message)
	for i := 0; i < len(vm); i++ {
		if i == 0 {
			if unicode.IsLetter(vm[0]) {
				upStr += string(vm[i] - 32)
			}
		} else {
			upStr += string(vm[i])
		}
	}
	return upStr
}
