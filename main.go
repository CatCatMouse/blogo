package main

import (
	"blog/internal/mod0"
	"fmt"
)

func main() {
	m := mod0.NewM("caicai", 18)
	fmt.Printf("name: %s,age=%d\n", m.GetName(), m.GetAge())
	m.SetAge(28)
	m.SetName("haha")
	fmt.Printf("name: %s,age=%d\n", m.GetName(), m.GetAge())
}
