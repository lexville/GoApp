package controllers

import (
	"GoApp/view"
	"log"
	"net/http"
)

// View contains the home view
type View struct {
	HomeView *view.View
}

// AddViewTemplate gets all the templates needed
// to render the view
func AddViewTemplate() *View {
	return &View{
		HomeView: view.ParseViewFiles(
			"base",
			"resources/views/welcome.gohtml",
		),
	}
}

// Home renders the view templates
func (v *View) Home(w http.ResponseWriter, r *http.Request) {
	if err := v.HomeView.Render(w, nil); err != nil {
		log.Panic("Unable to render the home view templates")
	}
}
