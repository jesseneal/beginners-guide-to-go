package main

func main() {
	var customerHeight int = 140
	customerAge := 18

	if customerHeight >= 150 || customerAge >= 18 {
		println("can access ride")
	} else if customerHeight >= 120 {
		println("customer can access children's rides")
	} else {
		println("Customer is too small")
	}
}
