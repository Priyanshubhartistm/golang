// struct in go = it is a basically coustm data structure hote h,
// jaise dusri programming ke andr classes bante h, or obj. banake usko use krte h, to kch isi trh se ham go ke andr class ki jhug "struct" use kar skte h

package main

import (
	"fmt"
	"time"
)

// multiple field ko v ek jhugh pe group krne ke liye v strsuct use kr skte h.

// ex:= ek order ko represent krne ke liye ham yaha pe ek struct bana skte h,
// order struct

type order struct {
	Id           int
	customerName string
	status       string
	amount       float64
}

func main() {
	myOrder := order{
		Id:           1,
		customerName: "John Doe",
		status:       "Pending",
		amount:       100.50,
	}

	myOrder.createdAt = time.Now()

	fmt.Println(myOrder.status)

	fmt.Println("order struct", myOrder)
	// order struct {Id int; customerName string; status string; amount float64; createdAt time.Time} {Id:1 customerName:John Doe status:Pending amount:100.5 createdAt:2024-06-20 18:29:45.123456 +0000 UTC m=+0.000123456}

	// struct ke andr function ni bna skte h, but struct k sath method bna skte h
}
