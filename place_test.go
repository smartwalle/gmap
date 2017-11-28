package gmap

import (
	"testing"
	"fmt"
)

func TestGoogleMap_NearbySearch(t *testing.T) {
	fmt.Println("=====NearbySearch=====")
	var p = NearbySearchParam{}
	p.Types = "store"
	p.Keyword = "vapor"
	p.Radius = 5000
	p.Latitude = -33.8670522
	p.Longitude = 151.1957362

	fmt.Print(client.NearbySearch(p))
}