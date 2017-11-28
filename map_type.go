package gmap

import "net/url"

type GoogleMapParam interface {
	Params() url.Values
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Viewport struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

type Geometry struct {
	Location *Location `json:"location"`
	Viewport *Viewport `json:"viewport"`
}

type Photo struct {
	Height int `json:"height"`
	Width int `json:"width"`
	HtmlAttributions []string `json:"html_attributions"`
	PhotoReference string `json:"photo_reference"`
}