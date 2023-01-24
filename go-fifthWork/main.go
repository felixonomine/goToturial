package main

import "fmt"

func main() {
	name := []string{"jarule", "make", "jay"}
	for index, value := range name {
		fmt.Printf("the value at index.v% is.v% \n", index, value)
	}

	i := 0
	for i < 5 {
		fmt.Println("The value of i is = ", i)
		i++

	}

	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++

	}
}
