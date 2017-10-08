package web

import (
	"html/template"
	"io"
)

func render(n string, wr io.Writer, data interface{}) error {
	t, err := template.ParseFiles("templates/" + n + ".html")
	if err != nil {
		return err
	}

	return t.Execute(wr, data)
}
