package estcequecest

import (
	//	"fmt"
	"strconv"
	//	"strings"
	"sort"
)

type timeline struct {
	messages map[int]string
	order    []int
}

// Get a message given minutes
func (tl *timeline) GetMessage(unit int) string {
	for _, time := range tl.order {
		if time <= unit {
			return tl.messages[time]
		}
	}
	return ""
}

// Pretty print
func (tl *timeline) String() string {
	txt := ""
	for _, time := range tl.order {
		unit := strconv.Itoa(time/60) + ":" + strconv.Itoa(time%60)
		txt += unit + ":" + tl.messages[time] + "\n"
	}
	return txt
}

func newTimeline(messages map[string]string) *timeline {
	sorted := make([]int, len(messages))
	intKeys := make(map[int]string, len(messages))
	index := 0
	// translate time code to int
	// create an array of sorted times
	for time, mess := range messages {
		hour, _ := strconv.Atoi(time[0:2])
		min, _ := strconv.Atoi(time[3:5])
		unit := hour*60 + min
		intKeys[unit] = mess
		sorted[index] = unit
		index++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	tl := new(timeline)
	tl.messages = intKeys
	tl.order = sorted
	return tl
}
