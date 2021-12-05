package main

import (
	"fmt"
)

func main() {
	planetSlice := []string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune"}

	for i := 0; i < len(planetSlice); i++ {
		fmt.Println(planetSlice[i])
	}

	for i := 0; i < len(planetSlice); {
		fmt.Println(planetSlice[i])
		i++
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	for i, v := range planetSlice {
		fmt.Println(i, v)
	}
}
