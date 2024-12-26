package main

import (
	"errors"
	"fmt"
)

var errorFraction = errors.New("Zero in the denominator")

func createFraction(intvar1, intvar2 int) (int, error) {
	if intvar2 == 0 {
		return 0, errorFraction
	}
	return intvar1 / intvar2, nil
}

func main() {
	intvar1, intvar2 := 14, 0
	output, error := createFraction(intvar1, intvar2)
	if error != nil {
		switch {
		case errors.Is(error, errorFraction):
			fmt.Println("Fraction with Zero Denominator")
		default:
			fmt.Printf("error in creating a fraction: %s\n", error)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", intvar1, intvar2, output)
}
