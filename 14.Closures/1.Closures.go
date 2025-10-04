// Closures

package main

// ek function k nam "counter" denge, fir y kch () recivie ni krega, but ek function return krega , or y jo func h, n y integer return krega,

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter() // yha pr counter() call hoga, or y ek function return krega, or oo function increment k variable m store ho jayega

	println(increment())

}
