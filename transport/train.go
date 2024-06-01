package transport

import "fmt"

type Train struct {
}

func NewTrain() *Train {
	return &Train{}
}

func (train *Train) BoardPassengers() {
	fmt.Println("Поїзд приймає пасажирів")
}

func (train *Train) DropPassengers() {
	fmt.Println("Поїзд висаджує пасажирів")
}
