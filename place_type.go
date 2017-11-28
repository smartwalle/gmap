package gmap

import (
	"fmt"
	"net/url"
)

type PlaceNearbySearchParam struct {
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

func (this PlaceNearbySearchParam) Params() url.Values {
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

type PlaceNearbySearchResults struct {
	HtmlAttributions []string          `json:"html_attributions"`
	Results          []*PlaceBasicInfo `json:"results"`
	NextPageToken    string            `json:"next_page_token"`
	Status           string            `json:"status"`
}

type PlaceBasicInfo struct {
	Id           string        `json:"id"`
	PlaceId      string        `json:"place_id"`
	Name         string        `json:"name"`
	Geometry     *Geometry     `json:"geometry"`
	Icon         string        `json:"icon"`
	Photos       []*Photo      `json:"photos"`
	Rating       float32       `json:"rating"`
	Reference    string        `json:"reference"`
	Scope        string        `json:"scope"`
	Types        []string      `json:"types"`
	Vicinity     string        `json:"vicinity"`
	OpeningHours *OpeningHours `json:"opening_hours"`
}

type PlaceDetailsParam struct {
	PlaceId    string
	Extensions string
	Language   string
}

func (this PlaceDetailsParam) Params() url.Values {
	var v = url.Values{}
	if len(this.PlaceId) > 0 {
		v.Add("placeid", this.PlaceId)
	}
	if len(this.Language) > 0 {
		v.Add("language", this.Language)
	}
	if len(this.Extensions) > 0 {
		v.Add("extensions", this.Extensions)
	}
	return v
}

type PlaceDetailsResults struct {
	HtmlAttributions []string      `json:"html_attributions"`
	Status           string        `json:"status"`
	Result           *PlaceDetails `json:"result"`
}

type PlaceDetails struct {
	*PlaceBasicInfo
	AddressComponents        []*AddressComponent `json:"address_components"`
	AdrAddress               string              `json:"adr_address"`
	FormattedAddress         string              `json:"formatted_address"`
	FormattedPhoneNumber     string              `json:"formatted_phone_number"`
	InternationalPhoneNumber string              `json:"international_phone_number"`
	Reviews                  []*Review           `json:"reviews"`
	URL                      string              `json:"url"`
	UTCOffset                int                 `json:"utc_offset"`
	Website                  string              `json:"website"`
}
