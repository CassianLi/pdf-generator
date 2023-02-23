package utils

import (
	"bytes"
	"html/template"
)

// ParseTemplate Parse a template to html
func ParseTemplate(templateFileName string, data interface{}) (page string, err error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
