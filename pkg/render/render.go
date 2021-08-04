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

//GOAL: We're not accessing the disk to get the template every time we load the page.
//Instead, we build a map that holds all of our templates, put that in an application wide site config
//and we can render our templates.

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the te,plate package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData data that want to be available on every page
func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td //templateData
}


// RenderTemplate takes response writer and string (name of template) pass- and read it to browser
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	// If in development mode(not production), don't use template cache, instead rebuild it on every request.
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error wirting template to brower", err)
	}

}

// CreateTemplateCache creates template as map, go through and pass our templates and return map of *.page.tmpl.html
func CreateTemplateCache() (map[string]*template.Template, error) {

	// Go to the templates folder, find everything in templates folder that's at the root level of my application
	//// and get all of the files that start with anything but definitely end with *.page.tmpl.html
	myCache := map[string]*template.Template{}

	// Go to template folder, find everything that start with *(anything) but end with .page.tmpl.html;
	pages, err := filepath.Glob("./templates/*.page.tmpl.html")
	if err != nil {
		return myCache, err
	}

	//For every page that I find, I am going to get the index and the first time through will be zero A
	//and then get the actual page itself, go to one because we start counting at zero
	//and four loops and the second one and the second value would be homePageTemplate
	//not just about page that template, it's turning the full path to that.
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
		// add ts (template set) the cache, use map we created.
		myCache[name] = ts
	}
	return myCache, nil
}

/* Conclusion:
So when I run this, my cache will actually have two entries:
	1. "about.page.tmpl.html"
	2. "home.page.tmpl.html
when look it up, it's going to give a fully parsed and ready to use template
It's actually a pointer to a template
When I need to render a template, I need to have access to this cache
*/
