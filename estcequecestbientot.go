package main

import (
	"estcequecestbientot/messageApp"
	"fmt"
)

func main() {
	//	a := messageApp.DoesItFit(5)
	b := messageApp.NewIntervalBit("4-9,20-68,*", 6)
	c := messageApp.NewIntervalBit("6,45-50", 6)
	fmt.Printf("Mes ib : %s | %s : %b %b\n", b, c, c.DoesItFit(7), c.DoesItFit(46))
}
