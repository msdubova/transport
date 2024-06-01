package transport

import "fmt"

type Bus struct {
}

func (bus Bus) BoardPassengers() {
	fmt.Println("Автобус приймає пасажирів")
}

func (bus Bus) DropPassengers() {
	fmt.Println("Автобус висаджує пасажирів")
}
