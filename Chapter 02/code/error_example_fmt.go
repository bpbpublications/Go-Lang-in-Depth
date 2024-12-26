package main

import "fmt"

func createFraction(intvar1, intvar2 int) {
	if intvar2 == 0 {
		return 0, fmt.Errorf("can't divide '%d' by zero", a)
	}
	return intvar1 / intvar2, nil
}

func main() {
	intvar1, intvar2 := 14, 0
	output, error := createFraction(intvar1, intvar2)

}
