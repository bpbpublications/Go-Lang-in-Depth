package main

type Displayer interface {
	display()
}

func main() {
	var varDisplay Displayer

	varDisplay.display()
}
