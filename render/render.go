package render

import (
	"bytes"
	"fmt"
	"github.com/Fietzorama/Bookings/pkg/config"
	"github.com/Fietzorama/Bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// Render has config, because we don't want to recreate the template cache
// every time we build a page.
var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData data that want to be available on every page
// taking the template data and returning it
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td //templateData
}

// Every single time we go to render any page on the site tis function is getting called
// it receives a response writer and it receives the name of the template i want to render

// Need to tell the Handlers to render the Templates in /templates
// RenderTemplate  renders templates using html/tempalte
// takes response writer and string (naming the template i want to render)
// and  pass- and read it to browser
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	// creating template cache when the application starts
	// when render a page, get the value from this application structure, our config.
	// I'm actually going to look in the configuration file and ask myself a simple question.
	//  declare the variable this is going to hold my template.
	var tc map[string]*template.Template

	// If in development mode(not production), don't use template cache,
	// instead rebuild it on every request
	// reference to main.go app.UseCache = false
	//if app.UseCache(config-bool) == true, then read the information from the template cache
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		//otherwise rebuild the templateCache
		tc, _ = CreateTemplateCache()
	}

	// if the template exists t  == true
	// Template Cache is active, we have template we want to use
	t, ok := tc[tmpl]
	// if t == FALSE, stop
	if !ok {
		log.Fatal("Could not get template from template_cache")
	}

	// t == true template is in the cache
	// Reading template from the cache and NOT from disk
	// creating bytes buffer "buf" to hold bytes that parsed.Template
	// has in memory into some bytes
	buf := new(bytes.Buffer)

	//
	td = AddDefaultData(td)

	// Similar to t,ok (not check for an error)
	// going to execute To the buffer, the value of that template
	// take this template that I have executed,
	// don't pass it any data and store the value in this buff variable.

	//We're not accessing the disk to get the template every time we load the page.
	//Instead, we build a map (tc) that holds all of our templates,
	//put that in an application wide site config
	// and we can render our templates.
	_ = t.Execute(buf, td)

	//  write "buf" to my response writer
	// and that writes everything to the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to Browser", err)
	}

}

// CreateTemplateCache creates template as map, go through
// and pass our templates and return map of *.page.tmpl.html
func CreateTemplateCache() (map[string]*template.Template, error) {

	// Go to the /templates folder, find everything end with *.page.tmpl.html
	// render them into a template combine them
	myCache := map[string]*template.Template{}

	// Go to template folder, find everything that start with *(anything) but end with .page.tmpl.html;
	pages, err := filepath.Glob("./templates/*.page.tmpl.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// print out current page
		fmt.Println("Page is currently", page)

		// input statement to create new function, a new variable,
		// call it functions and all it's going to do, it's equal to template.Func Map
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// Check: does this template match any layouts?
		// looking in templates directory for any file that ends .layout.tmpl.html
		matches, err := filepath.Glob("./templates/*.layout.tmpl.html")
		if err != nil {
			return myCache, err
		}
		// Check for errors at variable matches;
		//find at least one thing in there, then length of matches will be >0
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl.html")
			if err != nil {
				return myCache, err
			}
		}
		// add ts (template set) to the cache, use map we created.
		myCache[name] = ts
	}
	return myCache, nil
}
