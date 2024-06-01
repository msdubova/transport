package main

import (
	"transport/route"
	"transport/transport"
)

func main() {

	var bus = &transport.Bus{}
	var train = &transport.Train{}
	var plane = &transport.Plane{}

	var route = &route.Route{}

	route.AddTransport(bus)
	route.AddTransport(train)
	route.AddTransport(plane)
	route.AddTransport(bus)

	route.ShowTransport()
}
