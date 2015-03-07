// estcequecestbientot HTTP server
package main

import (
	"estcequecestbientot/estcequecest"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var app *estcequecest.App

func main() {
	app = estcequecest.NewApp("messages", "messages_")
	app.Load("pause")
	app.Load("apero")

	http.HandleFunc("/", index)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	host := os.Getenv("HOST")
	fmt.Println("App running on " + host + ":" + port)
	http.ListenAndServe(host+":"+port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	messages := app.GetMessages()
	t, _ := template.ParseFiles("templates/estcequecestbientot.html")
	t.Execute(w, messages)
}
