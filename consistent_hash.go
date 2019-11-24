package main

import "math"

const (
	KeyMax = math.MaxUint32
	KeyMin = uint32(0)
)

type ConsistentHash interface {
	Set([]byte)
	Get([]byte) ([]byte, bool)
}

type consistentHash struct {
	data map[uint32][]byte
}

func (c *consistentHash) Set(b []byte) {
	h := newMyHash(b)
	var index uint32
	h.Sum(&index)

	var key uint32
	for k := range c.data {
		if index < k {

		}
	}
}

func (c *consistentHash) Get([]byte) ([]byte, bool) {
	panic("implement me")
}

func newConsistentHash(count int) ConsistentHash {
	m := make(map[uint32][]byte)

	if count == 0 || count == 1 {
		m[0] = make([]byte, 0)
		return &consistentHash{data: m}
	}

	interval := KeyMax / (count - 1)

	for i := 1; i < count; i += interval {
		m[uint32(1)] = make([]byte, 0)
	}

	return &consistentHash{data: m}
}
