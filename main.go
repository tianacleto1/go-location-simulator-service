package main

import (
	"fmt"

	"com.anacleto.location-simulator/application/route"
)

func main() {

	route := route.Route{
		ID:       "1",
		ClientId: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()

	fmt.Println(stringJson[0])

}
