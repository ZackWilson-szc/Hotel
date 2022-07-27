package render

import (
	"Hotel/pkg/config"
	"Hotel/pkg/models"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.Appconfig

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	// when I discover to determine exactly what data I want to be available on every page, I can add here
	// but for now, this function is not useful
	return td
}

// NewTemplates sets the config for the template package
func NewTemplates(a *config.Appconfig) {
	app = a
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// create a template cache
	var tc map[string]*template.Template
	if app.UseCache { // 正式使用的时候，UseCache可以是True，不用每次加载，开发的时候改成false，每次刷新都会有改变
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from data
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	td = AddDefaultData(td)  // template data
	buf := new(bytes.Buffer) // these two lines are not mandatory, execute the value that I got from that map
	_ = t.Execute(buf, td)   // give me a clear view the value I got from the map, the error will appear if we parse it but can't execute
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)
	// get all files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl") // return all the files satisfied the pattern
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page) // base return the file name itself, so it will store in name
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// find the layout file
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl") // associate layout with tmpl
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts // after the parse
	}
	return myCache, nil
}

//// RenderTemplateOld renders templates using html/template
//// but this function will read all files from disk, and this is not efficient
//func RenderTemplateOld(w http.ResponseWriter, tmpl string) {
//	parseTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
//	err := parseTemplate.Execute(w, nil)
//	if err != nil {
//		fmt.Println("Error parsing template:", err)
//	}
//}
