package route

import (
	"fmt"
)

type Route struct {
	transportList []PublicTransport
}

func NewRoute() *Route {
	return &Route{}
}

type PublicTransport interface {
	BoardPassengers()
	DropPassengers()
}

func (route *Route) AddTransport(transport PublicTransport) {
	route.transportList = append(route.transportList, transport)
}

func (route *Route) ShowTransport() {
	fmt.Println("✅ ✅ ✅ Маршрут складається з: ")

	for _, item := range route.transportList {
		item.BoardPassengers()
		item.DropPassengers()
	}

	fmt.Println("Маршрут завершено  	🏁 🏁 🏁 🏁 🏁 🏁 ")
}
