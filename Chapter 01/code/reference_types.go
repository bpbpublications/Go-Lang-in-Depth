package main

import (
	"fmt"
	"reflect"
)

func changeSlice(slice []int) {
	reflect.ValueOf(&slice).Elem().SetLen(5)
	reflect.ValueOf(&slice).Elem().SetCap(5)
	fmt.Println("slice3>", len(slice), cap(slice))
}

func main() {
	slice := make([]int, 2, 8)
	fmt.Println("slice1>", len(slice), cap(slice))
	changeSlice(slice)
	fmt.Println("slice2>", len(slice), cap(slice))
}
