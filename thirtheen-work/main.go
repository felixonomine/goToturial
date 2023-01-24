package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}
func updateMenu(p map[string]float64) {
	p["coffee"] = 2.99
}

func main() {
	name := "onomine"
	name = updateName(updateName(name))
	fmt.Println(name)

	menu := map[string]float64{
		"pie":       5.99,
		"ice cream": 3.55,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
