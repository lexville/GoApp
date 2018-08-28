package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	viewsDirectory      = "resources/views/"
	templateExtension   = ".gohtml"
	layoutViewDirectory = "resources/views/layouts/"
)

// View contains a template and the layout
// to be used to render a view
type View struct {
	Template *template.Template
	Layout   string
}

// ParseViewFiles gets all the templates needed
// in a view and parses them. It then returns the
// view
func ParseViewFiles(layout string, files ...string) *View {
	files = append(files, parseLayoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal("Unable to parse the view files")
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// Render renders the view
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func parseLayoutFiles() []string {
	files, err := filepath.Glob(layoutViewDirectory + "*" + templateExtension)
	if err != nil {
		log.Fatal("Unable to get all the layout files")
	}
	return files
}
