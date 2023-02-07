package traffic

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func NewWeightedRoundRobin() Traffic {
	return &WeightedRoundRobin{
		mu:         sync.Mutex{},
		peers:      make([]*Peer, 0),
		counter:    0,
		peersCount: 0,
	}
}

// WeightedRoundRobin ...
type WeightedRoundRobin struct {
	mu         sync.Mutex
	peers      []*Peer
	peersCount int
	counter    int64
}

func (w *WeightedRoundRobin) String() string {
	str := fmt.Sprintf("\nTotal: %d\nCounter: %d\n", w.peersCount, w.counter)
	for i, p := range w.peers {
		str += fmt.Sprintf("Peer[%d]: %s (%d)\n", i+1, p.Name, p.Weight)
	}
	return str
}

func (w *WeightedRoundRobin) Reset() {
	w.counter = 0
	w.peersCount = 0
	w.peers = make([]*Peer, 0)
}

func (w *WeightedRoundRobin) Add(peer interface{}, weight int) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	for _, p := range w.peers {
		if p == peer {
			return fmt.Errorf("wrr: peer already exist")
		}
	}

	for i := 0; i < weight; i++ {
		w.peers = append(w.peers, &Peer{Name: peer, Weight: weight})
	}

	w.peersCount += weight

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(w.peersCount, func(i, j int) {
		w.peers[i], w.peers[j] = w.peers[j], w.peers[i]
	})

	return nil
}

func (w *WeightedRoundRobin) Next() interface{} {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.peersCount == 0 {
		return nil
	}

	p := w.peers[w.counter%int64(w.peersCount)]
	w.counter++

	if w.counter >= int64(w.peersCount) {
		w.counter = 0
	}

	return p.Name
}
