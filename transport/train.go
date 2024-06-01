package transport

import "fmt"

type Train struct {
}

func (train Train) BoardPassengers() {
	fmt.Println("Поїзд приймає пасажирів")
}

func (train Train) DropPassengers() {
	fmt.Println("Поїзд висаджує пасажирів")
}
