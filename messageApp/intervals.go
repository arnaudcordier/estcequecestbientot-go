package messageApp

import (
	"fmt"
	"strings"
	"time"
)

type Interval struct {
	intervals []*IntervalBit
}

// Create a new Interval
func NewInterval(s string) *Interval {
	i := new(Interval)
	i.intervals = createInterval(s)
	return i
}

func (i Interval) GetMessage(t time.Time) (string, bool) {
//	fmt.Printf("%d \n", times)
	if fits, ok := i.doesItFit(t); ok {
		return fmt.Sprintf("It fits : %s", fits), true
	}
	return "", false
}

func (i Interval) doesItFit(t time.Time) ([]Bits, bool) {
	// Convert time to array of integer, weekdays from 1 to 7
	times := []int{t.Year(), int(t.Month()), t.Day(), (int(t.Weekday())+6)%7+1, t.Hour(), t.Minute()}
	if len(times) == len(i.intervals) {
		fitingInterval := make([]Bits, len(times))
		for i, bit := range i.intervals {
			if itFits, ok := bit.DoesItFit(times[i]); ok {
				fitingInterval[i] = itFits
			//	fmt.Printf("%d fits in '%s'\n", times[i], itFits)
			} else {
				fmt.Printf("%d does not fit in '%s'\n", times[i], bit)
				return []Bits{}, false
			}
		}
		return fitingInterval, true
	}
	return []Bits{}, false
}

// split interval string and create IntervalBit from it
func createInterval(s string) []*IntervalBit {
	intervals := strings.Split(s, " ")
	bits := make([]*IntervalBit, len(intervals))
	for i, val := range intervals {
		bits[i] = NewIntervalBit(val, i+1)
	}
	return bits
}

// pretty print for an Interval
func (i *Interval) String() string {
	s := ""
	for _, val := range(i.intervals) {
		s += val.String() + " "
	}
	return s
}