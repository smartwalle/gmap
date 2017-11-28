package gmap

import (
	"testing"
	"fmt"
)

func TestGoogleMap_PlaceNearbySearch(t *testing.T) {
	fmt.Println("=====PlaceNearbySearch=====")
	var p = PlaceNearbySearchParam{}
	p.Types = K_PLACE_TYPE_STORE
	p.Keyword = "vapor"
	p.Radius = 5000
	p.Latitude = -33.8670522
	p.Longitude = 151.1957362

	var r, _ = client.PlaceNearbySearch(p)
	if r == nil {
		t.Fatal("PlaceNearbySearch 获取数据异常")
	}
	if r.Status != "OK" {
		t.Fatal("PlaceNearbySearch 返回数据状态错误")
	}
}

func TestGoogleMap_PlaceTextSearch(t *testing.T) {
	fmt.Println("=====PlaceTextSearch=====")
	var p = PlaceTextSearchParam{}
	p.Types = K_PLACE_TYPE_STORE
	p.Query = "apple in china"

	var r, _ = client.PlaceTextSearch(p)
	if r == nil {
		t.Fatal("PlaceTextSearch 获取数据异常")
	}
	if r.Status != "OK" {
		t.Fatal("PlaceTextSearch 返回数据状态错误")
	}
}

func TestGoogleMap_PlaceRadarSearch(t *testing.T) {
	fmt.Println("=====PlaceRadarSearch=====")
	var p = PlaceRadarSearchParam{}
	p.Types = K_PLACE_TYPE_STORE
	p.Radius = 5000
	p.Latitude = -33.8670522
	p.Longitude = 151.1957362

	var r, _ = client.PlaceRadarSearch(p)
	if r == nil {
		t.Fatal("PlaceRadarSearch 获取数据异常")
	}
	if r.Status != "OK" {
		t.Fatal("PlaceRadarSearch 返回数据状态错误")
	}
}

func TestGoogleMap_PlaceDetails(t *testing.T) {
	fmt.Println("=====PlaceDetails=====")
	var p = PlaceDetailsParam{}
	p.PlaceId = "ChIJN1t_tDeuEmsRUsoyG83frY4"
	var r, _ = client.PlaceDetails(p)
	if r == nil {
		t.Fatal("PlaceDetails 获取数据异常")
	}
	if r.Status != "OK" {
		t.Fatal("PlaceDetails 返回数据状态错误")
	}
}