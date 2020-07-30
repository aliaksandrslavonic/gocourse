package main

import "fmt"

type plant interface {
	grow()
}

type violet struct {
}

func (v *violet) grow() {
	fmt.Println("I will grow like small bush")
}

type rose struct {
}

func (r *rose) grow() {
	fmt.Println("I will grow like big bush")
}

func main() {
	var v, r plant = &violet{}, &rose{}
	v.grow()
	r.grow()
}
