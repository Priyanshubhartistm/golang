// kise functions ko dusre fun. ke andr as a arguments v pass kr skte ho,  function ke andr function pass krna
// y fir aap kise fun. ko variables ko v assign kr skte ho
// y fir ek fun. se ek new function v rturn kr skte ho

package main

import "fmt"

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}

}

func main() {
	// fn := func(a int) int {
	// 	return 2
	// }

	fn := processIt()
	fn(6)

	fmt.Println(fn(6)) // 4
}
