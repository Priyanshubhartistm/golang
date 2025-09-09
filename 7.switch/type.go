// type switch

whoAmI := func(i interface{}) {

	switch t := i.(type) {

	case int:
		fmt.Println("its an integer")

	case string:
		fmt.Println("its an string")

	}

case bool:
	fmt.Println("")

default:
	fmt.Println("")
}