package main

import "fmt"

func main() {
	h := newConsistentHash(4)
	key := []byte("today")
	value := []byte("hello")
	h.Add(key, value)
	fmt.Printf("add key: %s, value: %s\n", string(key), string(value))
	out, exist, hashKey := h.Get(key)
	fmt.Printf("value: %s, exist: %t, hash key: %d\n", string(out), exist, hashKey)
	err := h.Remove(0)
	if nil != err {
		fmt.Printf("remove index 1, data already exist")
		return
	}
	out, exist, hashKey = h.Get(key)
	fmt.Printf("remove 1 value: %s, exist: %t, hash key: %d\n", string(out), exist, hashKey)
}
