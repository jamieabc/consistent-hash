package main

import (
	"fmt"
	"math"
)

const (
	KeyMax = math.MaxUint32
)

type ConsistentHash interface {
	Add([]byte, []byte)
	Get([]byte) ([]byte, bool, uint32)
	Remove(int) error
}

type consistentHash struct {
	data map[uint32][]byte
	keys []uint32
}

// add key value
func (c *consistentHash) Add(key []byte, value []byte) {
	h := hashKey(c, key)
	c.data[h] = value
}

func hashKey(c *consistentHash, key []byte) uint32 {
	h := newMyHash(key)
	var hashNum uint32
	h.Sum(&hashNum)

	for i, k := range c.keys {
		if i == len(c.keys)-1 {
			break
		}
		if k < hashNum && c.keys[i+1] > hashNum {
			return c.keys[i+1]
		}
	}
	return c.keys[0]
}

// get value of key
func (c *consistentHash) Get(key []byte) ([]byte, bool, uint32) {
	k := hashKey(c, key)
	if len(c.data[k]) == 0 {
		return []byte{}, false, 0
	}
	return c.data[k], true, k
}

func (c *consistentHash) Remove(index int) error {
	storedKey := c.keys[index]
	var storedValue []byte

	if len(c.data[storedKey]) != 0 {
		storedValue = c.data[storedKey]
	}

	removeByIndex(index, c)

	err := redistribute(c, storedKey, storedValue)
	return err
}

func removeByIndex(index int, c *consistentHash) {
	// remove from key
	k := c.keys[index]
	if index == 0 {
		c.keys = c.keys[1:]
	} else {
		c.keys = append(c.keys[0:index], c.keys[index+1:]...)
	}

	// remove from map
	delete(c.data, k)
}

func redistribute(c *consistentHash, key uint32, value []byte) error {
	target := c.keys[0]
	for i, k := range c.keys {
		if k < key && c.keys[i+1] > key {
			target = c.keys[i+1]
			break
		}
	}

	if len(c.data[target]) != 0 {
		return fmt.Errorf("data already exist")
	} else {
		c.data[target] = value
	}
	return nil
}

func newConsistentHash(count int) ConsistentHash {
	m := make(map[uint32][]byte)

	if count == 0 || count == 1 {
		m[0] = make([]byte, 0)
		return &consistentHash{
			data: m,
			keys: []uint32{0},
		}
	}

	interval := KeyMax / count
	keys := make([]uint32, count)
	keys[0] = uint32(interval / 2)
	m[keys[0]] = make([]byte, 0)

	for i := 1; i < count; i++ {
		key := uint32(interval * (2*i + 1) / 2)
		m[key] = make([]byte, 0)
		keys[i] = key
	}

	return &consistentHash{
		data: m,
		keys: keys,
	}
}
