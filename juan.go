package main

import "fmt"

var times float32 = 0.1

const BASE = 525600

func main() {
	fmt.Println("beautiful world.  start juan")
	juan()
}

func juan() {
	times += 1
	fmt.Printf("you are in juan year:%f \n", times/BASE)
	juan()
}
