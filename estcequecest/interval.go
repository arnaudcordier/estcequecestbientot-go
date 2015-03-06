package estcequecest

import (
	"fmt"
	"strconv"
	"strings"
)

type interval struct {
	sInterval string
	name      string
	boundary  bit
	bits      []bit
	duration  int
}

type bit struct {
	all bool
	min int
	max int
}

// Definition of types of IntervalBit, used with pos to determine the kind of interval
var (
	bitsDef = []interval{
		interval{"2000-3000", "Year", bit{false, 2000, 3000}, []bit{}, 365},
		interval{"1-12", "Month", bit{false, 1, 12}, []bit{}, 31},
		interval{"1-31", "Day of month", bit{false, 1, 31}, []bit{}, 24},
		interval{"1-7", "Day of week", bit{false, 1, 7}, []bit{}, 24},
		interval{"0-23", "Hour", bit{false, 0, 23}, []bit{}, 60},
		interval{"0-59", "Minute", bit{false, 0, 59}, []bit{}, 1},
	}
)

func (i *interval) doesItFit(time int) (bit, bool) {
	for _, b := range i.bits {
		if b.all == true || (b.min <= time && time <= b.max) {
			return b, true
		}
	}
	return bit{}, false
}

// create an interval using an intervalString and a position
func newInterval(s string, pos int) *interval {
	pos = pos - 1
	i := new(interval)
	i.sInterval = s
	i.name = bitsDef[pos].name
	i.boundary = bitsDef[pos].boundary
	i.duration = bitsDef[pos].duration
	i.bits = createBits(s, i.boundary)
	return i
}

// create an array of bit from a string like '6,45-50,*'
// uses bit "boundary" to validate the interval
func createBits(s string, boundary bit) []bit {
	bits := strings.Split(s, ",")
	allBits := make([]bit, len(bits))
	nValidBits := 0
	for _, val := range bits {
		thisBit := bit{false, 0, 0}
		if val == "*" {
			thisBit.all = true
			allBits[nValidBits] = thisBit
			nValidBits++
			continue
		}
		parts := strings.Split(val, "-")
		thisBit.min, _ = strconv.Atoi(strings.Trim(parts[0], " "))
		if len(parts) == 1 {
			thisBit.max = thisBit.min
		} else {
			thisBit.max, _ = strconv.Atoi(strings.Trim(parts[1], " "))
		}
		// valid if min and max fit in boundary
		if boundary.min <= thisBit.min && thisBit.min <= boundary.max &&
			boundary.min <= thisBit.max && thisBit.max <= boundary.max {
			allBits[nValidBits] = thisBit
			nValidBits++
		} else {
			fmt.Printf("Error in interval creation: %s could not fit in boundary %s\n", val, boundary)
		}
	}
	return allBits[0:nValidBits]
}

// Pretty print for a Bits
func (b bit) String() string {
	if b.all == true {
		return "*"
	}
	s := fmt.Sprintf("%d-%d", b.min, b.max)
	return s
}

// Pretty print for an IntervalBit
func (i *interval) String() string {
	s := "<" + i.name + ": "
	for _, val := range i.bits {
		s += val.String() + ","
	}
	s = s[0:len(s)-1] + ">"
	return s
}
