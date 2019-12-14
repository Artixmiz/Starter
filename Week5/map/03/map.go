package main

import "fmt"

func main() {
	elements := make(map[string]map[string]string)
	elements["H"] = map[string]string{"name": "Hydrogen", "state": "gas"}
	elements["He"] = map[string]string{"name": "Helium", "state": "gas"}
	fmt.Println(elements)
	fmt.Println(elements["He"]["state"])
}
