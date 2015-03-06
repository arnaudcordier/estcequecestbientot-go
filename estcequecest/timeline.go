package estcequecest

import (
	"fmt"
	"sort"
	"strconv"
)

type timeline struct {
	messages map[int]string
	order    []int
}

// Get a message given minutes
func (tl *timeline) GetMessage(unit int) string {
	for _, minutes := range tl.order {
		if minutes <= unit {
			return tl.messages[minutes]
		}
	}
	return ""
}

// Pretty print
func (tl *timeline) String() string {
	txt := ""
	n := len(tl.order)
	for i := n - 1; i >= 0; i -= 1 {
		minutes := tl.order[i]
		txt += fmt.Sprintf("%02d:%02d : '%s'\n", minutes/60, minutes%60, tl.messages[minutes])
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
		minutes := hour*60 + min
		intKeys[minutes] = mess
		sorted[index] = minutes
		index++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	tl := new(timeline)
	tl.messages = intKeys
	tl.order = sorted
	return tl
}
