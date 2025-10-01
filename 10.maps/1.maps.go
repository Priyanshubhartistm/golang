// Maps  = Associated data structure hota h.
// if you come to another language thn you knoe to, jo Hash table, js ke andr obj, python ke andr dictionary hote h,
// to usi trh k data structure chahiya to apko maps use krna hoga
// maps are reference types
// maps are unordered collections of key-value
// keys are unique
// keys are of same type
// values are of same
// zero value of map is nil
package main

import "fmt"

func main() {

	// creating map

	// map ko sbse phle nam dte h, kch v nam de skte h jaisea(m,n)
	// Then Make function use krenge jo slices ke liye use ki thi hamne, = make()
	// Then keyword likhna h,
	// map ke bad dena h, square bracket[]
	// or ise bracket ke andr type likhna h, map ke "key" k type denge.
	// or ise bracket ke bhar map ke "value" k type denge.

	m := make(map[string]string)

	// setting an element

	m["name"] = "golang"
	m["area"] = "backend"

	// get an element and

	// m likhne ke bad
	// bracket deke usme key dete h.
	fmt.Println(m["name"], m["area"])

	// IMP:- if key does not exists in the map then it returns zero value
	fmt.Println(m["phone"]) // o/p = empty
}
