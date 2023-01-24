package main

import "fmt"

func main() {
	names := []string{"jow", "peter", "john", "mark", "love"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			println("breaking at the pos", index)
			break
		}
		fmt.Printf("the value at position %v is %v\n", index, value)
	}
}
