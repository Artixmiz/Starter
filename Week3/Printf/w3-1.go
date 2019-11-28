package main

import "fmt"

func main() {
	fmt.Printf("My name is %s, Iam %d, years old\n", "goku", 18)

	fmt.Printf("type = %T \n", 3.14159265359)
	fmt.Printf("PI = %f \n", 3.14159265359)
	fmt.Printf("PI = %2.2f \n", 3.14159265359)
	fmt.Printf("PI = %9.f \n", 3.14159265359)
	fmt.Printf("PI = %-9.f \n", 3.14159265359)
	fmt.Printf("PI = %09.f \n", 3.14159265359)
	fmt.Printf("PI = %09.2f \n", 3.14159265359)
	fmt.Printf("PI = %E \n", 3.14159265359)

	fmt.Printf("1 > 2 = %t \n", 1 > 2)
	fmt.Printf("10 (base 2) = %b \n", 10)
	fmt.Printf("10 (base 8) = %o \n", 10)
	fmt.Printf("10 (base 10) = %d \n", 10)
	fmt.Printf("10 (base 16) = %x \n", 10)
}
