package templateparse

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type templateLoader struct {
	rootDirectory   string
	globalTemplates *template.Template
}

// TemplateLoader global state for loading template
var TemplateLoader = templateLoader{
	rootDirectory:   "./webserver",
	globalTemplates: template.Must(findAndParseTemplates("./webserver", nil)),
}

func (tl *templateLoader) render(w http.ResponseWriter, template string, data interface{}) error {

	w.Header().Add("Content-type", "text/html; charset=utf-8")
	return tl.globalTemplates.ExecuteTemplate(w, template, data)
}

// RenderTemplate for html template rendering
func RenderTemplate(w http.ResponseWriter, template string, data interface{}) error {
	return TemplateLoader.render(w, template, data)
}

//var globalTemplates *template.Template = template.Must(FindAndParseTemplates("./webserver", nil))

func findAndParseTemplates(rootDir string, funcMap template.FuncMap) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			t := root.New(name).Funcs(funcMap)
			t, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	return root, err
}
