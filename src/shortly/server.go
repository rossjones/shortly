package main

import (
	"fmt"
    "html/template"
	"net/http"
	"path"
)

var template_root string
var templates *template.Template

func root_handler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "index.html", map[string]interface{}{"Title": "Test"})
}

func api_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love APIS")
}

func app_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love APP")
}

func renderTemplate(w http.ResponseWriter, tmpl string, ctx map[string]interface{}) {
    // When in debug, we will load the templates each time so we can modify them
    // without restarting
    if configuration.App.Debug {
        t, _ := template.ParseFiles(path.Join(template_root,tmpl ))
        t.Execute(w, ctx)
        return
    }

    err := templates.ExecuteTemplate(w, tmpl, ctx)
    if err != nil {
        // Log the problem
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func run_server() error {
    template_root = path.Join(configuration.Templates.Path, "templates/")
    if ! configuration.App.Debug {
        templates = template.Must(template.ParseFiles(path.Join(template_root, "base.html"),
                                                      path.Join(template_root, "index.html")))
    }

    svr := fmt.Sprintf("%s:%d", configuration.Server.Host, configuration.Server.Port)
	fmt.Printf("Shortly is listening on %s\n", svr)

	http.HandleFunc("/", root_handler)
	http.HandleFunc("/api/", api_handler)
	http.HandleFunc("/app/", app_handler)

	// Static assets
	media := path.Join(configuration.Templates.Path, "assets/")
	http.Handle("/media/",
		http.StripPrefix("/media/", http.FileServer(http.Dir(media))))

	http.ListenAndServe(svr, nil)

	return nil
}
