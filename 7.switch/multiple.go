// multiple switch

package main

import (
	"fmt"
	"time"
)

func main() {

	// Multiple switch

	switch time.Now().weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")

	default:
		fmt.Println("it's workday")
	}

}
