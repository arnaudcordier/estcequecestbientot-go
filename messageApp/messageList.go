package messageApp

import (
//	"fmt"
	"strconv"
//	"strings"
	"sort"
)

type MessageList struct {
	messages map[int]string
	order []int
}

// Get a message given minutes
func (ml *MessageList) GetMessage(unit int) string {
	for _, time := range ml.order {
		if time <= unit {
			return ml.messages[time]
		}
	}
	return ""
}

// Pretty print
func (ml *MessageList) String() string {
	txt := ""
	for _, time := range ml.order {
		unit := strconv.Itoa(time/60) + ":" + strconv.Itoa(time%60)
		txt += unit + ":" + ml.messages[time] + "\n"
	}
	return txt
}

func NewMessageList(messages *map[string]string) *MessageList {
	sorted := make([]int, len(*messages))
	intKeys := make(map[int]string, len(*messages))
	index := 0
	// translate time code to int
	// create an array of sorted times
	for time, mess := range *messages {
		hour, _ := strconv.Atoi(time[0:2])
		min, _ := strconv.Atoi(time[3:5])
		unit := hour * 60 + min
		intKeys[unit] = mess
		sorted[index] = unit
		index++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	ml := new(MessageList)
	ml.messages = intKeys
	ml.order = sorted
	return ml
}