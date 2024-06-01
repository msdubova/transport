package transport

import "fmt"

type Plane struct {
}

func NewPlane() *Plane {
	return &Plane{}
}

func (plane *Plane) BoardPassengers() {
	fmt.Println("Літак приймає пасажирів")
}

func (plane *Plane) DropPassengers() {
	fmt.Println("Літак висаджує пасажирів")
}
