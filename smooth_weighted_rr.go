package traffic

import (
	"fmt"
	"sync"
)

/*
SW (Smooth Weighted) is a struct that contains weighted items and provides methods to select a weighted item.
It is used for the smooth weighted round-robin balancing algorithm. This algorithm is implemented in Nginx:
https://github.com/phusion/nginx/commit/27e94984486058d73157038f7950a0a36ecc6e35.
http://hg.nginx.org/nginx/rev/c90801720a0c

Algorithm is as follows: on each peer selection we increase current_weight
of each eligible peer by its weight, select peer with greatest current_weight
and reduce its current_weight by total number of weight points distributed
among peers.

In case of { 5, 1, 1 } weights this gives the following sequence of
current_weight's: (a, a, b, a, c, a, a)
*/

func NewSmoothWeightedRoundRobin() Traffic {
	return &SmoothWeightedRoundRobin{
		mu:         sync.Mutex{},
		peers:      make([]*Peer, 0),
		peersCount: 0,
	}
}

type SmoothWeightedRoundRobin struct {
	mu         sync.Mutex
	peers      []*Peer
	peersCount int
}

func (w *SmoothWeightedRoundRobin) String() string {
	str := fmt.Sprintf("\nCounter: %d\n", w.peersCount)
	for i, p := range w.peers {
		str += fmt.Sprintf("Peer[%d]: %s (%d)\n", i+1, p.Name, p.Weight)
	}
	return str
}

// Add a weighted server.
func (w *SmoothWeightedRoundRobin) Add(peer interface{}, weight int) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	for _, p := range w.peers {
		if p == peer {
			return fmt.Errorf("wrr: peer already exist")
		}
	}

	w.peers = append(w.peers, &Peer{Name: peer, Weight: weight, EffectiveWeight: weight})
	w.peersCount++

	return nil
}

// Next returns selected server.
// https://github.com/phusion/nginx/commit/27e94984486058d73157038f7950a0a36ecc6e35
func (w *SmoothWeightedRoundRobin) Next() interface{} {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.peersCount == 0 {
		return nil
	}

	if w.peersCount == 1 {
		return w.peers[0].Name
	}

	total := 0
	var best *Peer

	for i := 0; i < w.peersCount; i++ {
		p := w.peers[i]

		if p == nil {
			continue
		}

		p.CurrentWeight += p.EffectiveWeight
		total += p.EffectiveWeight
		if p.EffectiveWeight < p.Weight {
			p.EffectiveWeight++
		}

		if best == nil ||
			p.CurrentWeight > best.CurrentWeight {
			best = p
		}

	}

	if best == nil {
		return nil
	}

	best.CurrentWeight -= total

	return best.Name
}

func (w *SmoothWeightedRoundRobin) Reset() {
	for _, p := range w.peers {
		p.EffectiveWeight = p.Weight
		p.CurrentWeight = 0
	}
}
