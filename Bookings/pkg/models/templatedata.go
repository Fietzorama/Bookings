package models

// send data from handlers to templates (pages)
// "models" store all models that used in the application
// every time rendering a template, I can manually construct all
// of data that I need to pass to that template
// you want to take a template, data, whatever template that is passed by a handler,
// and then add to it data that you want available on every page of your site
// includes database models and it includes the template data model.

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string      //sending n strings
	IntMap    map[string]int         //sending n int
	FloatMap  map[string]float32     //sending n float
	Data      map[string]interface{} // sending n Datatypes
	CSRFToken string                 //Cross Site Request Forgery to construct post forms and handle the forum post
	// hidden field in form == string random numbers, change when some goes to a page
	Flash   string // Msg to the user (error msg, warning etc.)
	Warning string // Msg to the user (error msg, warning etc.)
	Error   string // Msg to the user (error msg, warning etc.)
}
