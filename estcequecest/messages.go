package estcequecest

import (
	// 	"fmt"
	"strings"
	// 	"time"
)

type messages struct {
	cron     []*interval
	timeline *timeline
}

// Create a new Interval
func newMessages(data messagesData) *messages {
	m := new(messages)
	m.cron = createIntervals(data.Interval)
	m.timeline = newTimeline(data.Timeline)
	return m
}

// pretty print for an messages
func (m *messages) String() string {
	s := ""
	for _, interval := range m.cron {
		s += interval.String() + " "
	}
	s += "\n" + m.timeline.String()
	return s
}

// func (i *Interval) GetMessage(t time.Time) (string, bool) {
// 	// Convert time to array of integer, weekdays from 1 to 7
// 	times := []int{t.Year(), int(t.Month()), t.Day(), (int(t.Weekday())+6)%7 + 1, t.Hour(), t.Minute()}
// 	if _, ok := i.doesItFit(times); ok {
// 		now := times[4]*60 + times[5]
// 		return i.messages.GetMessage(now), true
// 	}
// 	return "", false
// }
//
// // does the given time fits in the intervalBits
// func (i *Interval) doesItFit(times []int) ([]Bits, bool) {
// 	// check each intervalBits if it fits with given time
// 	if len(times) == len(i.intervals) {
// 		fitingInterval := make([]Bits, len(times))
// 		for i, bit := range i.intervals {
// 			if itFits, ok := bit.DoesItFit(times[i]); ok {
// 				fitingInterval[i] = itFits
// 				//	fmt.Printf("%d fits in '%s'\n", times[i], itFits)
// 			} else {
// 				fmt.Printf("%d does not fit in '%s'\n", times[i], bit)
// 				return []Bits{}, false
// 			}
// 		}
// 		return fitingInterval, true
// 	}
// 	return []Bits{}, false
// }
//
// split interval string and create Intervals from it
func createIntervals(s string) []*interval {
	splited := strings.Split(s, " ")
	intervals := make([]*interval, len(splited))
	for i, val := range splited {
		intervals[i] = newInterval(val, i+1)
	}
	return intervals
}
