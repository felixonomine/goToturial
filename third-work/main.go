package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var ages = []int{6, 7, 9, 2, 1, 0, 3, 5}
	var name = [3]string{"apple\n mango\n fruit"}
	name[1] = "orrange"
	greeting := "hello there onomine!\n"
	ages[0] = 100

	fmt.Println(ages, len(ages))
	fmt.Println(name, len(name))
	fmt.Println(strings.ToUpper(name[1]))
	fmt.Printf(strings.ToUpper(greeting))
	fmt.Print(greeting)
	sort.Ints(ages)
	fmt.Println(ages)
	sort.Strings(name[:])
	fmt.Println(name)
}
