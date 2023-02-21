package main

import "fmt"

func main() {
	book := map[string]float64{
		"maths":   7.99,
		"physis":  6.44,
		"biology": 9.55,
		"english": 5.66,
	}
	fmt.Println(book)
	for k, v := range book {
		fmt.Println(k, "", v)
		fmt.Println(k, "-", v)
		for index, book := range book {
			fmt.Println("all of them:\n", index, book)
			book++
		}

	}
	price := map[string]int{
		"beans": 7,
		"yam":   8,
		"soup":  5,
		"gari":  9,
	}
	fmt.Println(price)
	//for k, v := range price {
	//	fmt.Println("the price\n", k, ":", v)
	//}
}
