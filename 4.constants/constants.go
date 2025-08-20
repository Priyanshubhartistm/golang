package main

import "fmt"

func main() {
	// const name = "golang"
	// const age = 30

	// shorthand syntax
	// :=

	// multiple constants
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)

}
