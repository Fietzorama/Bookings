package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string      //sending n strings
	IntMap    map[string]int         //sending n int
	FloatMap  map[string]float32     //sending n float
	Data      map[string]interface{} // sending n Datatypes
	CSRFToken string                 //Cross Site Request Forgery to construct post forms and handle the forum post
									// hidden field in form is long string random numbers, it change when some goes to a page
	Flash     string                 // Msg to the user (error msg, warning etc.)
	Warning   string                // Msg to the user (error msg, warning etc.)
	Error     string                // Msg to the user (error msg, warning etc.)
}

