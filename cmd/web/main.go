package main

// package always live in their own directory, all the files for a package must exist in the same directory
import (
	"Hotel/pkg/config"
	"Hotel/pkg/handlers"
	"Hotel/pkg/render"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.Appconfig
var session *scs.SessionManager

// main is the application function
func main() {
	// init global config
	var app config.Appconfig

	// session
	app.InProduction = false // also can be used to substitute UseCache
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session // for other part to use this
	// get the template cache from the config

	tc, err := render.CreateTemplateCache() // use render to initialize template cache
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc    // init template cache
	app.UseCache = false      // this can switch between developer mood and user mood
	render.NewTemplates(&app) // give back render the access of app

	// handlers part
	repo := handlers.NewRepo(&app) // create the repository variable
	handlers.NewHandlers(repo)     // pass it back to handler

	//// old route
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	err = srv.ListenAndServe()
	log.Fatal(err)
	// _ = http.ListenAndServe(portNumber, nil)

}

// basic connection about web request and response
//func main() {
//	// handle with web, the first argument is url
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		n, err := fmt.Fprintf(w, "Hello World!") // Fprintf require a response writer to write to
//		if err != nil {
//			fmt.Println(err)
//		}
//		// :n will automatically format these return types to what ever these functions return back
//		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
//		// SpringF allow me to print any type and return as a string
//	})
//	// start a web server that listens for requests in go
//	_ = http.ListenAndServe(":8080", nil)
//	// listen on the TCP network and then call Serve with Handler to handle request
//}
