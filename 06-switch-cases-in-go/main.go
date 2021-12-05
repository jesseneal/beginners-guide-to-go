package main

import "runtime"

func main() {

	switch os := runtime.GOOS; os {
	case "darwin":
		println("os x")
	case "linux":
		println("linux machine")
	default:
		println("something else")
	}

}
