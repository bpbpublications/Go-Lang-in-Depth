package main

import "fmt"

func createFraction(intvar1, intvar2 int) (int, error) {
	if intvar2 == 0 {
		return 0, fmt.Errorf("A fraction cannot have a denominator as zero and '%d' as numerator", intvar2)
	}
	return intvar1 / intvar2, nil
}

func main() {

	intvar1, intvar2 := 14, 0
	output, error := createFraction(intvar1, intvar2)

	fmt.Printf("Error is %s and Output is %d\n", error.Error(), output)

}
