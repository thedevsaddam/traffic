package traffic

import (
	"testing"
)

func TestNewWeightedRoundRobin(t *testing.T) {
	rr := NewWeightedRoundRobin()
	rr.Add("a", 5)
	rr.Add("b", 2)
	rr.Add("c", 3)

	m := map[interface{}]int{}
	for i := 0; i < 100; i++ {
		m[rr.Next()]++
	}
	if m["a"] != 50 && m["b"] != 20 && m["c"] != 30 {
		t.Fail()
	}
}
