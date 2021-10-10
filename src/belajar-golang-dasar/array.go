package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Yusuf"
	names[1] = "Nur"
	names[2] = "Wahid"

	for i := 0; i < 3; i++ {
		fmt.Println(names[i])
	}

	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values[2])

	var Values [10]int

	// Panjang array
	fmt.Println(len(names))
	fmt.Println(len(values))
	fmt.Println(len(Values))
}
