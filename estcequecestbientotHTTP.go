// estcequecestbientot HTTP server
package main

import (
	"estcequecestbientot/estcequecest"
	"html/template"
	"net/http"
)

var app *estcequecest.App

func main() {
	app = estcequecest.NewApp("messages", "messages_")
	app.Load("pause")
	app.Load("apero")

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	messages := app.GetMessages()
	t, _ := template.ParseFiles("templates/estcequecestbientot.html")
	t.Execute(w, messages)
}
