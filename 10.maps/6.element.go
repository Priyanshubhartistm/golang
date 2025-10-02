// how to check if a key exists in a map
// how to check basically elements exists in a map or not

package main

func main() {

	m := map[string]int{"price": 40, "phone": 100}

	_, ok := m["price"] // price key exists krta h ki ni, to ok true hoga, agr ni krta to false hoga

	if ok {
		println("all ok")
	} else {
		println("not ok")
	}

}
