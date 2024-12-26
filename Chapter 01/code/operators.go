package main

import "fmt"

func main() {
	var1 := 46
	var2 := 87

	//Arithmetic Operators

	// Addition Operator
	output1 := var1 + var2
	fmt.Printf("Result of var1 + var2 = %d", output1)

	// Subtraction Operator
	output2 := var1 - var2
	fmt.Printf("\nResult of var1 - var2 = %d", output2)

	// Multiplication Operator
	output3 := var1 * var2
	fmt.Printf("\nResult of var1 * var2 = %d", output3)

	// Division Operator
	output4 := var1 / var2
	fmt.Printf("\nResult of var1 / var2 = %d", output4)

	// Modulus Operator
	output5 := var1 % var2
	fmt.Printf("\nResult of var1 %% var2 = %d", output5)

	//Relational Operators

	// ‘=='(Equal To Operator)
	output11 := var1 == var2
	fmt.Println(output11)

	// ‘!='(Not Equal To Operator)
	output12 := var1 != var2
	fmt.Println(output12)

	// ‘<‘(Less Than Operator)
	output13 := var1 < var2
	fmt.Println(output13)

	// ‘>'(Greater Than Operator)
	output14 := var1 > var2
	fmt.Println(output14)

	// ‘>='(Greater Than Equal To Operator)
	output15 := var1 >= var2
	fmt.Println(output15)

	// ‘<='(Less Than Equal To Operator)
	output16 := var1 <= var2
	fmt.Println(output16)

	//Logical Operators

	if var1 != var2 && var1 <= var2 {
		fmt.Println("True")
	}

	if var1 != var2 || var1 <= var2 {
		fmt.Println("True")
	}

	if !(var1 == var2) {
		fmt.Println("True")
	}

	//Bitwise Operators

	// & (bitwise AND Operator)
	output21 := var1 & var2
	fmt.Printf("Result of var1 & var2 = %d", output21)

	// | (bitwise OR Operator)
	output22 := var1 | var2
	fmt.Printf("\nResult of var1 | var2 = %d", output22)

	// ^ (bitwise XOR Operator)
	output23 := var1 ^ var2
	fmt.Printf("\nResult of var1 ^ var2 = %d", output23)

	// << (left shift Operator)
	output24 := var1 << 1
	fmt.Printf("\nResult of var1 << 1 = %d", output24)

	// >> (right shift Operator)
	output25 := var1 >> 1
	fmt.Printf("\nResult of var1 >> 1 = %d", output25)

	// &^ (AND NOT Operator)
	output26 := var1 &^ var2
	fmt.Printf("\nResult of var1 &^ var2 = %d", output26)

	//Assignment Operators

	// “=”(Simple Assignment Operator)
	var1 = var2
	fmt.Println(var1)

	// “+=”(Add Assignment Operator)
	var1 += var2
	fmt.Println(var1)

	//“-=”(Subtract Assignment Operator)
	var1 -= var2
	fmt.Println(var1)

	// “*=”(Multiply Assignment Operator)
	var1 *= var2
	fmt.Println(var1)

	// “/=”(Division Assignment Operator)
	var1 /= var2
	fmt.Println(var1)

	// “%=”(Modulus Assignment Operator)
	var1 %= var2
	fmt.Println(var1)

}
