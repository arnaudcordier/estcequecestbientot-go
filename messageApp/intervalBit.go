package messageApp

import (
	"fmt"
	"strings"
	"strconv"
)

type IntervalBit struct {
	sInterval	string
	name	string
	boundary Bits
	intervals	[]Bits
	duration	int
}

type Bits struct {
	all	bool
	min	int
	max	int
}

// Definition of types of IntervalBit, used with pos to determine the kind of interval
var (
	bitsDef = []IntervalBit{
		IntervalBit{"2000-3000",	"Year",			Bits{false, 2000, 3000},	[]Bits{},	365},
		IntervalBit{"1-12",		"Month",			Bits{false, 1, 12},		[]Bits{},	31},
		IntervalBit{"1-31",		"Day of month",	Bits{false, 1, 31},		[]Bits{},	24},
		IntervalBit{"1-7",		"Day of week",		Bits{false, 1, 7},			[]Bits{},	24},
		IntervalBit{"0-23",		"Hour",			Bits{false, 0, 23},		[]Bits{},	60},
		IntervalBit{"0-59",		"Minute",			Bits{false, 0, 59},		[]Bits{},	1},
	}
)

func (ib *IntervalBit) DoesItFit(time int) bool {
	for _, bit := range ib.intervals {
		if bit.all == true || (bit.min <= time && time <= bit.max) {
			return true
		}
	}
	return false
}

// create an IntervalBit using an intervalString and a position
func NewIntervalBit(s string, pos int) *IntervalBit {
	ib := new(IntervalBit)
	pos = pos-1
	ib.sInterval = s
	ib.name = bitsDef[pos].name
	ib.boundary = bitsDef[pos].boundary
	ib.duration = bitsDef[pos].duration
	ib.createIntervals(s)
	return ib
}

// create an array of intervals from a string like '6,45-50,*'
func (ib *IntervalBit) createIntervals(s string) {
	bits := strings.Split(s, ",")
	allBits := make([]Bits, len(bits))
	nValidBits := 0
	for _, val := range bits {
		thisBit := Bits{false, 0, 0}
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
		if ib.boundary.min<=thisBit.min && thisBit.min<=ib.boundary.max &&
		   ib.boundary.min<=thisBit.max && thisBit.max<=ib.boundary.max {
			allBits[nValidBits] = thisBit
			nValidBits++
		} else {
			fmt.Printf("%s: %s could not fit in boundary %s\n", ib, val, ib.boundary)
		}
	}
	ib.intervals = allBits[0:nValidBits]
}

// Pretty print for a Bits
func (b Bits) String() string {
	if b.all == true {
		return "*"
	}
	s :=  fmt.Sprintf("%d-%d", b.min, b.max)
	return s
}

// Pretty print for an IntervalBit
func (ib *IntervalBit) String() string {
	s := "<" + ib.name +": "
	for _, val := range ib.intervals {
		s += val.String() + ","
	}
	s = s[0:len(s)-1] + ">"
	return s
}

