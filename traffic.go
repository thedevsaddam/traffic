package traffic

import (
	"fmt"
)

// Traffic represents a round robin algorithm contract
type Traffic interface {
	Add(peer interface{}, weight int) error
	Next() interface{}
	Reset()
	fmt.Stringer
}

type Peer struct {
	Name            interface{}
	Weight          int
	CurrentWeight   int
	EffectiveWeight int
}
