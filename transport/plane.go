package transport

import "fmt"

type Plane struct {
}

func (plane Plane) BoardPassengers() {
	fmt.Println("Літак приймає пасажирів")
}

func (plane Plane) DropPassengers() {
	fmt.Println("Літак висаджує пасажирів")
}
