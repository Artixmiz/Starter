package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	h := elements["H"]
	fmt.Println(h)
	n, found := elements["N"]
	fmt.Println(n)
	fmt.Println(found)
}
