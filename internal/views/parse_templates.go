package views

import (
	"html/template"
	"net/http"
)

func ParseTmpl(w http.ResponseWriter, data interface{}, filenames ...string) error {

	path := "./web/templates/"

	for key, value := range filenames {
		filenames[key] = path + value + ".tmpl"
	}

	// for Template Name
	templateName := filenames[0][len(path):]

	templates, err := template.New(templateName).ParseFiles(filenames...)
	if err != nil {
		return err
	}

	if err = templates.Execute(w, data); err != nil {
		return err
	}

	return nil
}
