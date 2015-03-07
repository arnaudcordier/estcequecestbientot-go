package main

import (
	"estcequecestbientot/estcequecest"
	"fmt"
)

func main() {
	app := estcequecest.NewApp("messages", "messages_")
	// 	app.Load("object")
	app.Load("noon")
	app.Load("apero")
	app.Load("pause")
	list, list2 := app.List()
	fmt.Printf("%v %v\n", list, list2)
	app.Unload("apero")
	list, list2 = app.List()
	fmt.Printf("%v %v\n", list, list2)

	messages := app.GetMessages()
	fmt.Printf("Messages : %v\n", messages)
}
