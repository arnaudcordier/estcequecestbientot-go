// estcequecestbientot HTTP server
package main

import (
	"estcequecestbientot/estcequecest"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var app *estcequecest.App

func main() {
	app = estcequecest.NewApp("messages", "messages_")

	http.HandleFunc("/", index)
	http.HandleFunc("/load/", load)
	http.HandleFunc("/unload/", unload)
	http.HandleFunc("/list", list)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	host := os.Getenv("HOST")
	fmt.Println("Est-ce-que c'est running on " + host + ":" + port)

	http.ListenAndServe(host+":"+port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	messages := app.GetMessages()
	t, _ := template.ParseFiles("templates/estcequecestbientot.html")
	t.Execute(w, &messages)
}

func load(w http.ResponseWriter, r *http.Request) {
	name := strings.Replace(r.URL.Path, "/load/", "", 1)
	if len(name) > 0 {
		app.Load(name)
	}
	http.Redirect(w, r, "/", 303)
}

func unload(w http.ResponseWriter, r *http.Request) {
	name := strings.Replace(r.URL.Path, "/unload/", "", 1)
	if len(name) > 0 {
		app.Unload(name)
	}
	http.Redirect(w, r, "/", 303)
}

func list(w http.ResponseWriter, r *http.Request) {
	notloaded, loaded := app.List()
	t, _ := template.ParseFiles("templates/list.html")
	t.Execute(w, map[string]*[]string{"notloaded": &notloaded, "loaded": &loaded})
}
