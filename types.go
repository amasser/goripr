package goripr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

var (
	// GlobalLowerBoundary is the smallest value possible that indicates the lower boundary of the IP set
	GlobalLowerBoundary = &IPAttributes{ID: "-inf", LowerBound: false, UpperBound: true}

	//GlobalUpperBoundary is the biggest value possible that indicates the upper boundary of the IP set
	GlobalUpperBoundary = &IPAttributes{ID: "+inf", LowerBound: true, UpperBound: false}
)

// IPAttributes is the composite result type of anything requesting an IP
type IPAttributes struct {
	ID         string
	IP         net.IP
	Reason     string
	LowerBound bool
	UpperBound bool
}

// Equal tests if two attribute instances are equal
func (ia *IPAttributes) Equal(other *IPAttributes) bool {

	return ia.ID != "" &&
		ia.ID == other.ID &&
		ia.LowerBound == other.LowerBound &&
		ia.UpperBound == other.UpperBound &&
		ia.IP.Equal(other.IP) &&
		ia.Reason == other.Reason
}

// IPInt64 returns the IP's int64 value
func (ia *IPAttributes) IPInt64() int64 {
	bigInt, bits := IPToInt(ia.IP)
	// only working with IPv4
	if bits > IPv4Bits {
		panic("contained value is bigger than IPv4")
	}
	return bigInt.Int64()
}

// EqualIP returns true if the IPs of both are equal
func (ia *IPAttributes) EqualIP(other *IPAttributes) bool {
	return ia.IP.Equal(other.IP)
}

// IsInfBoundary returns true if ia is either the GlobalUpperBoundary or the GlobalLowerBoundary
func (ia *IPAttributes) IsInfBoundary() bool {
	return ia.Equal(GlobalLowerBoundary) || ia.Equal(GlobalUpperBoundary)
}

// IsSingleBoundary returns true
func (ia *IPAttributes) IsSingleBoundary() bool {
	if ia.LowerBound != ia.UpperBound {
		return true
	} else if ia.LowerBound && ia.UpperBound {
		return false
	}
	panic(errors.New("did not expect both boundaries to be false"))
}

func (ia *IPAttributes) String() string {
	b, err := json.Marshal(ia)
	if err != nil {
		panic(fmt.Errorf("failed to marshal IPAttributes: %w", err))
	}
	return string(b)
}

// IPRangeAttributes maps an IP range to the Reason string
type IPRangeAttributes struct {
	Range  string
	Reason string // Ban reason
}
