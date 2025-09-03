package main

import "fmt"

func main() {
	for i := 0; i <= 3; i++ {
		// break = loop ko stop krne ke liye use hota h,
		if i == 2 {
			continue

			// 		continue = continue keyword h, o current iteration ko skip kr deta h, jaise:- (2) print ni hoga direct, (0,1,3) print ho jayega.
			// continue keyword loop ke andar hi use hota h.
			// isse hum loop ke andar se bahar nahi nikalte, bas current iteration ko skip karte hain.

		}
		fmt.Println(i)

	}

}
