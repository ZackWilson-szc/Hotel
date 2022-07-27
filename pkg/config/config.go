package config

// this is global variables for all files in this application
import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

type Appconfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
