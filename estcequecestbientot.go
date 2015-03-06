package main

import (
	"estcequecestbientot/estcequecest"
	"fmt"
)

func main() {
	app := estcequecest.NewApp("messages", "messages_")
	app.Load("object")
	app.Load("noon")
	app.Load("apero")
	app.Load("pause")
	fmt.Printf("%s", app)
	messages := app.GetMessages()
	fmt.Printf("Messages : %v\n", messages)
}
