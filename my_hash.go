package main

const (
	seed = 1234567890
)

type MyHash interface {
	Write([]byte) int
	Sum(*uint32)
}

type myHash struct {
	data []byte
}

func (h *myHash) Write(b []byte) int {
	h.data = b
	return len(h.data)
}

func (h *myHash) Sum(out *uint32) {
	var tmp uint32
	for _, d := range h.data {
		tmp *= seed
		tmp ^= uint32(d)
	}
	*out = tmp
}

func newMyHash(b []byte) MyHash {
	return &myHash{
		data: b,
	}
}
