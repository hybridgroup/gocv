package gocv

import (
	"sync"
)

var CGOCallback = callback{
	fns: make(map[int]func(...interface{})),
}

type callback struct {
	mu    sync.Mutex
	index int
	fns   map[int]func(...interface{})
}

func (n *callback) Register(fn func(args ...interface{})) int {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.index++
	for n.fns[n.index] != nil {
		n.index++
	}
	n.fns[n.index] = fn
	return n.index
}

func (n *callback) Lookup(i int) func(...interface{}) {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.fns[i]
}

func (n *callback) Unregister(i int) {
	n.mu.Lock()
	defer n.mu.Unlock()
	delete(n.fns, i)
}
