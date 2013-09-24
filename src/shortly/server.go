package main

import (
	"fmt"
    "github.com/flosch/pongo"
	"net/http"
	"path"
)

var templates = make(map[string]*pongo.Template)
var tpl_root string

func root_handler(w http.ResponseWriter, r *http.Request) {
    renderTemplate("index.html", pongo.Context{"Title": "test"}, w)
}

func api_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love APIS")
}

func app_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love APP")
}

func renderTemplate(tmpl string, ctx pongo.Context, w http.ResponseWriter) {
    // When in debug, we will load the templates each time so we can modify them
    // without restarting
    if configuration.App.Debug {
        t, err := pongo.FromFile(path.Join(tpl_root, tmpl), nil)
        err = t.ExecuteRW(w, &ctx)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    err := templates[tmpl].ExecuteRW(w, &ctx)
    if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func run_server() error {
    tpl_root = path.Join(configuration.Templates.Path, "templates/")
    templates["index.html"] = pongo.Must(pongo.FromFile(path.Join(tpl_root, "index.html"), nil))


    svr := fmt.Sprintf("%s:%d", configuration.Server.Bind, configuration.Server.Port)
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
