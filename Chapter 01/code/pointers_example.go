package main

import "fmt"

func main() {

	hexa_var1 := 0xCC
	hexa_var2 := 0x8F

	fmt.Printf("variable type of hexa_var1 is %T\n", hexa_var1)
	fmt.Printf("hexa_var1's value in hexadecimal is %X\n", hexa_var1)
	fmt.Printf("hexa_var1's value in decimal is %v\n", hexa_var1)

	fmt.Printf("variable type of hexa_var2 is %T\n", hexa_var2)
	fmt.Printf("hexa_var2's value in hexadecimal is %X\n", hexa_var2)
	fmt.Printf("hexa_var2's value in decimal is %v\n", hexa_var2)

}
