package main

import (
	"bytes"
	"html/template"
)

func RenderTemplate(record *Record) (string, error) {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		return "", err
	}
	var renderedTemplate bytes.Buffer
	err = tmpl.Execute(&renderedTemplate, record)
	if err != nil {
		return "", err
	}
	return renderedTemplate.String(), nil
}