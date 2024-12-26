package main

import "fmt"

type IPhoto interface {
	getTitle() string
}

type Painting struct {
	title string
}

func (paint *Painting) getTitle() string {
	return paint.title
}

type YellowFrame struct {
	photo IPhoto
}

func (yframe *YellowFrame) getTitle() string {
	return yframe.photo.getTitle()
}

func main() {

	painting := new(Painting)

	painting.title = "Mona Lisa"

	yframe := &YellowFrame{photo: painting}

	fmt.Println("The tile of the yellowFramed painting is ", yframe.getTitle())

}
