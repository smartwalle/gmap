package gmap

import "net/url"

type GoogleMapParam interface {
	Params() url.Values
}

type Status string

const (
	K_STATUS_OK               Status = "OK"
	K_STATUS_ZERO_RESULTS     Status = "ZERO_RESULTS"
	K_STATUS_OVER_QUERY_LIMIT Status = "OVER_QUERY_LIMIT"
	K_STATUS_REQUEST_DENIED   Status = "REQUEST_DENIED"
	K_STATUS_INVALID_REQUEST  Status = "INVALID_REQUEST"
)

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
	Height           int      `json:"height"`
	Width            int      `json:"width"`
	HtmlAttributions []string `json:"html_attributions"`
	PhotoReference   string   `json:"photo_reference"`
}

type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type OpeningHours struct {
	OpenNow     bool                  `json:"open_now"`
	Periods     []*OpeningHoursPeriod `json:"periods"`
	WeekdayText []string              `json:"weekday_text"`
}

type OpeningHoursPeriod struct {
	Close *OpeningHoursPeriodTime `json:"close"`
	Open  *OpeningHoursPeriodTime `json:"open"`
}

type OpeningHoursPeriodTime struct {
	Day  int    `json:"day"`
	Time string `json:"time"`
}

type Review struct {
	AuthorName              string  `json:"author_name"`
	AuthorURL               string  `json:"author_url"`
	Language                string  `json:"language"`
	ProfilePhotoURL         string  `json:"profile_photo_url"`
	Rating                  float32 `json:"rating"`
	RelativeTimeDescription string  `json:"relative_time_description"`
	Time                    int64   `json:"time"`
}
