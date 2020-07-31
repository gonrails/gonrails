package helper

import (
	"fmt"
	"os"
	"text/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"bold":       bold,
		"red":        red,
		"capitalize": capitalize,
	}
}

func Tmpl(text string, data interface{}) {
	output := os.Stdout
	t := template.New("HowTo").Funcs(FuncMap())
	template.Must(t.Parse(text))
	err := t.Execute(output, data)
	PanicError(err)
}

func Black(str string) string {
	return textColor(textBlack, str)
}

func Red(str string) string {
	return textColor(textRed, str)
}
func Yellow(str string) string {
	return textColor(textYellow, str)
}
func Green(str string) string {
	return textColor(textGreen, str)
}
func Cyan(str string) string {
	return textColor(textCyan, str)
}
func Blue(str string) string {
	return textColor(textBlue, str)
}
func Purple(str string) string {
	return textColor(textPurple, str)
}
func White(str string) string {
	return textColor(textWhite, str)
}

func textColor(color int, str string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, str)
}
