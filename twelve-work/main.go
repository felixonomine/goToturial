package main

import "fmt"

func main() {
	menu := map[string]float32{
		"soup":  4.99,
		"pie":   7.99,
		"salad": 6.99,
		"garri": 3.55,
	}

	fmt.Println(menu["pie"])
	fmt.Println(menu["garri"])
	for l, v := range menu {
		fmt.Println(l, "=", v)
	}
	phoneNumber := map[int]string{
		+2348067155766: "evuarhere",
		+2347067155766: "onomine",
		+2348054870981: "erhiga",
		+2347054110981: "obara",
	}
	for k, p := range phoneNumber {
		fmt.Println(k, "=", p)
	}
}
