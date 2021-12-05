package main

import (
	"fmt"
)

func main() {
	planets := [8]string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune"}
	fmt.Println(planets)

	var planetsArray [8]string
	planetsArray[0] = "mercury"

	fmt.Println(planetsArray)
	planetSlice := []string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune"}
	fmt.Println(planetSlice)

	var planetSliceVerbose []string
	planetSliceVerbose = append(planetSliceVerbose, "mercury")
	fmt.Println(planetSliceVerbose)
}
