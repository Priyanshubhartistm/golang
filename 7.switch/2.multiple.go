// golang me  different types k switch hote h
// 1. simple switch
// 2. multiple switch
// 3. type switch

package main

import (
	"fmt"
	"time"
)

func main() {

	// Multiple switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")

	default:
		fmt.Println("it's workday")
	}

}
