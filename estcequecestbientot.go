package main

import (
	"estcequecestbientot/estcequecest"
	"fmt"
	// 	"time"
)

func main() {
	// intervalBits
	// 	a := messageApp.NewIntervalBit("4-9,20-68,*,59", 6)
	// 	b := messageApp.NewIntervalBit("6,45-50", 6)
	// 	f1, _ := b.DoesItFit(7)
	// 	f2, _ := b.DoesItFit(46)
	// 	f3, _ := a.DoesItFit(10)
	// 	fmt.Printf("Mes ib : %s | %s : %s %s %s \n", a, b, f1, f2, f3)
	// 	// Intervals
	// 	myMessage := map[string]string{"10:30": "message1", "15:45": "message2", "13:12": "message3", "09:12": "message4"}
	// 	myInterval := "* * 13-30 3-7 09-20,23 *"
	// 	d := messageApp.NewInterval(myInterval, &myMessage)
	// 	fmt.Printf("Interval %s \n", d)
	// 	t := time.Now()
	// 	f, _ := d.GetMessage(t)
	// 	fmt.Printf("Ça fit %s\n", f)
	// 	// Message List
	// 	ml := messageApp.NewMessageList(&myMessage)
	// 	fmt.Printf("%s \n", ml)
	app := estcequecest.NewApp("messages", "messages_")
	app.Load("object")
	fmt.Printf("%s \n", app)
}
