package models

// TemplateData holds data that will be sent to the templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[int]int
	FloatMap  map[string]float32
	DataMap   map[string]interface{}
	CSRFToken string //cross site request forgery token -
	Flash     string
	Warning   string
	Error     string
}
