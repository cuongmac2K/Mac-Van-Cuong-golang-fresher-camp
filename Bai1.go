package main

import (
	"fmt"
)

func crawl(s []int, k chan int) {
	for i := 0; i < len(s); i++ {
		k <- s[i]
	}
}

func main() {
	map1 := make(chan int, 250)
	map2 := make(chan int, 250)
	map3 := make(chan int, 250)
	map4 := make(chan int, 250)
	s := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = i + 1
	}
	go crawl(s[:250], map1)
	go crawl(s[250:500], map2)
	go crawl(s[500:750], map3)
	go crawl(s[750:1000], map4)

	fmt.Printf("\n channel One: \n")
	for i := 0; i < 250; i++ {
		fmt.Print(<-map1, "; ")
	}
	fmt.Printf("\n channel Two: \n")
	for i := 0; i < 250; i++ {
		fmt.Print(<-map2, "; ")
	}
	fmt.Printf("\n channel Three: \n")
	for i := 0; i < 250; i++ {
		fmt.Print(<-map3, "; ")
	}
	fmt.Printf("\n channel Four: \n")
	for i := 0; i < 250; i++ {
		fmt.Print(<-map4, "; ")
	}

}
