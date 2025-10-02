// 1. index accessing in range for slices

package main

func main() {
	// nums := []int{2, 3, 4, 5}

	// for i, num := range nums {
	// 	println(i, num) // 0 2 1 3 2 4 3 5
	// }

	// 2. index accessing in range for maps

	m := map[string]string{"name": "john", "course": "golang"}

	for k, v := range m {
		println(k, v) // name john course golang

	}

	// 3. single value accessing in range for maps
	for k := range m {
		println(k) // name course
	}

	// 4. string iteration using range
	for i, ch := range "hello" {
		//rintln(i, ch) // 0 104 1 101 2 108 3 108 4 111

		//agr apko har ek character ko print krna h to apko type conversion krna padega
		println(i, string(ch)) // 0 h 1 e 2 l 3 l 4 o

	}

}
