package main

import "fmt"

func main() {
	for i := 0; i <= 3; i++ {
		// break = loop ko stop krne ke liye use hota h,
		if i == 2 {
			continue

			// 		continue = continue keyword h, o current iteration ko skip
		}
		fmt.Println(i)

	}

}
