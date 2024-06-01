package main

import (
	"transport/route"
	"transport/transport"
)

func main() {

	var bus = transport.NewBus()
	var train = transport.NewTrain()
	var plane = transport.NewPlane()
	var coach = transport.NewBus()

	var route = route.NewRoute()

	route.AddTransport(bus)
	route.AddTransport(train)
	route.AddTransport(plane)
	route.AddTransport(coach)

	route.ShowTransport()
}
