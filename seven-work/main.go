package main

import "fmt"

func sayGreeeting(n string) {
	fmt.Printf("Good morning %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}
func cycleName(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {
	sayGreeeting("Evuarhere")
	sayGreeeting("onomine")
	sayBye("Jarule")

	cycleName([]string{"oke", "tega", "kevwe", "kessi"}, sayGreeeting)
	cycleName([]string{"oke", "tega", "kevwe", "kessi"}, sayBye)
}
