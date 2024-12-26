package main

import "fmt"

func getAverage(varArray [3]int) float32 {

	total := 0
	for i := 0; i < len(varArray); i++ {

		total = total + varArray[i]
	}

	return float32(total) / float32(len(varArray))

}

func main() {

	var varArray [3]int

	varArray = [3]int{2, 3, 4}

	var average float32
	average = getAverage(varArray)

	fmt.Printf("Average = %f\n", average)
}
