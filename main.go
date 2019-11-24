package main

import "fmt"

func main() {
	h := newMyHash("cat")
	var out uint32
	h.Sum(&out)
	fmt.Printf("%d\n", out)
}
