package gmap

import (
	"fmt"
	"net/url"
)

type NearbySearchParam struct {
	Latitude  float64
	Longitude float64
	Radius    int
	RankBy    string

	Keyword  string
	Language string
	//MinPrice int
	//MaxPrice int
	Name      string
	OpenNow   bool
	Types     string
	PageToken string
}

func (this NearbySearchParam) Params() url.Values {
	var v = url.Values{}
	v.Add("location", fmt.Sprint(this.Latitude, ",", this.Longitude))
	if len(this.RankBy) > 0 {
		v.Add("rankby", this.RankBy)
	} else if this.Radius > 0 {
		v.Add("radius", fmt.Sprint(this.Radius))
	}
	if len(this.Keyword) > 0 {
		v.Add("keyword", this.Keyword)
	}
	if len(this.Language) > 0 {
		v.Add("language", this.Language)
	}
	if len(this.Name) > 0 {
		v.Add("name", this.Name)
	}
	if this.OpenNow {
		v.Add("opennow", fmt.Sprint(this.OpenNow))
	}
	if len(this.Types) > 0 {
		v.Add("types", this.Types)
	}
	if len(this.PageToken) > 0 {
		v.Add("pageToken", this.PageToken)
	}
	return v
}

type NearbySearchResults struct {
	HtmlAttributions []string               `json:"html_attributions"`
	Results          []*NearbySearchResults `json:"results"`
	NextPageToken    string                 `json:"next_page_token"`
	Status           string                 `json:"status"`
}

type NearbySearchResultsItem struct {
	Geometry  *Geometry `json:"geometry"`
	Icon      string    `json:"icon"`
	Id        string    `json:"id"`
	Photos    []*Photo  `json:"photos"`
	PlaceId   string    `json:"place_id"`
	Rating    int       `json:"rating"`
	Reference string    `json:"reference"`
	Scope     string    `json:"scope"`
	Types     []string  `json:"types"`
	Vicinity  string    `json:"vicinity"`
}
