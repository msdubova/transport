package route

import (
	"fmt"
	"transport/transport"
)

type Route struct {
	transportList []transport.PublicTransport
}

func (route *Route) AddTransport(transport transport.PublicTransport) {
	route.transportList = append(route.transportList, transport)
}

func (route *Route) ShowTransport() {
	fmt.Println("Маршрут складається з: ")

	for _, item := range route.transportList {
		item.BoardPassengers()
		item.DropPassengers()
	}
}
