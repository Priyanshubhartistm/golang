package main

import "fmt"

func main() {
	// constants = aise variables h, jinka value change ni hota h.
	// const keyword se declare krte h.
	// const ka value hum kabhi change ni kr skte h.
	// const ka value hum program ke andar kahi bhi use kr skte h.

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
