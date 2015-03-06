package main

import (
	"estcequecestbientot/estcequecest"
	"fmt"
)

func main() {
	app := estcequecest.NewApp("messages", "messages_")
	app.Load("object")
	fmt.Printf("%s", app)
	messages := app.GetMessages()
	fmt.Printf("Messages : %v\n", messages)
}
