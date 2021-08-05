package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// avoid loading the template cache every single time
// we display a page on the site by setting configurations
// set the configuration such that once I have this template set,
// and I never want to load it again until the application restarts
// or I issue a command that says load the template cache
// create a package here is this configuration file might be accessed
// from each other all over the place any part of my application.
// be very careful to ensure that this configuration
// file doesn't import anything other than what it absolutely has to

// Changes in two places like InProduction in Sessions, put it here and make it global
// or make something from one package publi to other see Session

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger //write log files to everywhere i want
	InProduction  bool
	Session       *scs.SessionManager
}
