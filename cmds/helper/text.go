package helper

import (
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
