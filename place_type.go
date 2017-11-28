package gmap

import (
	"fmt"
	"net/url"
)

type PlaceType string

const (
	K_PLACE_TYPE_ACCOUNTING              PlaceType = "accounting"
	K_PLACE_TYPE_AIRPORT                 PlaceType = "airport"
	K_PLACE_TYPE_AMUSEMENT_PARK          PlaceType = "amusement_park"
	K_PLACE_TYPE_AQUARIUM                PlaceType = "aquarium"
	K_PLACE_TYPE_ART_GALLERY             PlaceType = "art_gallery"
	K_PLACE_TYPE_ATM                     PlaceType = "atm"
	K_PLACE_TYPE_BAKERY                  PlaceType = "bakery"
	K_PLACE_TYPE_BANK                    PlaceType = "bank"
	K_PLACE_TYPE_BAR                     PlaceType = "bar"
	K_PLACE_TYPE_BEAUTY_SALON            PlaceType = "beauty_salon"
	K_PLACE_TYPE_BICYCLE_STORE           PlaceType = "bicycle_store"
	K_PLACE_TYPE_BOOK_STORE              PlaceType = "book_store"
	K_PLACE_TYPE_BOWLING_ALLEY           PlaceType = "bowling_alley"
	K_PLACE_TYPE_BUS_STATION             PlaceType = "bus_station"
	K_PLACE_TYPE_CAFE                    PlaceType = "cafe"
	K_PLACE_TYPE_CAMPGROUND              PlaceType = "campground"
	K_PLACE_TYPE_CAR_DEALER              PlaceType = "car_dealer"
	K_PLACE_TYPE_CAR_RENTAL              PlaceType = "car_rental"
	K_PLACE_TYPE_CAR_REPAIR              PlaceType = "car_repair"
	K_PLACE_TYPE_CAR_WASH                PlaceType = "car_wash"
	K_PLACE_TYPE_CASINO                  PlaceType = "casino"
	K_PLACE_TYPE_CEMETERY                PlaceType = "cemetery"
	K_PLACE_TYPE_CHURCH                  PlaceType = "church"
	K_PLACE_TYPE_CITY_HALL               PlaceType = "city_hall"
	K_PLACE_TYPE_CLOTHING_STORE          PlaceType = "clothing_store"
	K_PLACE_TYPE_CONVENIENCE_STORE       PlaceType = "convenience_store"
	K_PLACE_TYPE_COURTHOUSE              PlaceType = "courthouse"
	K_PLACE_TYPE_DENTIST                 PlaceType = "dentist"
	K_PLACE_TYPE_DEPARTMENT_STORE        PlaceType = "department_store"
	K_PLACE_TYPE_DOCTOR                  PlaceType = "doctor"
	K_PLACE_TYPE_ELECTRICIAN             PlaceType = "electrician"
	K_PLACE_TYPE_ELECTRONICS_STORE       PlaceType = "electronics_store"
	K_PLACE_TYPE_EMBASSY                 PlaceType = "embassy"
	K_PLACE_TYPE_ESTABLISHMENT           PlaceType = "establishment"
	K_PLACE_TYPE_FINANCE                 PlaceType = "finance"
	K_PLACE_TYPE_FIRE_STATION            PlaceType = "fire_station"
	K_PLACE_TYPE_FLORIST                 PlaceType = "florist"
	K_PLACE_TYPE_FOOD                    PlaceType = "food"
	K_PLACE_TYPE_FUNERAL_HOME            PlaceType = "funeral_home"
	K_PLACE_TYPE_FURNITURE_STORE         PlaceType = "furniture_store"
	K_PLACE_TYPE_GAS_STATION             PlaceType = "gas_station"
	K_PLACE_TYPE_GENERAL_CONTRACTOR      PlaceType = "general_contractor"
	K_PLACE_TYPE_GROCERY_OR_SUPERMARKET  PlaceType = "grocery_or_supermarket"
	K_PLACE_TYPE_GYM                     PlaceType = "gym"
	K_PLACE_TYPE_HAIR_CARE               PlaceType = "hair_care"
	K_PLACE_TYPE_HARDWARE_STORE          PlaceType = "hardware_store"
	K_PLACE_TYPE_HEALTH                  PlaceType = "health"
	K_PLACE_TYPE_HINDU_TEMPLE            PlaceType = "hindu_temple"
	K_PLACE_TYPE_HOME_GOODS_STORE        PlaceType = "home_goods_store"
	K_PLACE_TYPE_HOSPITAL                PlaceType = "hospital"
	K_PLACE_TYPE_INSURANCE_AGENCY        PlaceType = "insurance_agency"
	K_PLACE_TYPE_JEWELRY_STORE           PlaceType = "jewelry_store"
	K_PLACE_TYPE_LAUNDRY                 PlaceType = "laundry"
	K_PLACE_TYPE_LAWYER                  PlaceType = "lawyer"
	K_PLACE_TYPE_LIBRARY                 PlaceType = "library"
	K_PLACE_TYPE_LIQUOR_STORE            PlaceType = "liquor_store"
	K_PLACE_TYPE_LOCAL_GOVERNMENT_OFFICE PlaceType = "local_government_office"
	K_PLACE_TYPE_LOCKSMITH               PlaceType = "locksmith"
	K_PLACE_TYPE_LODGING                 PlaceType = "lodging"
	K_PLACE_TYPE_MEAL_DELIVERY           PlaceType = "meal_delivery"
	K_PLACE_TYPE_MEAL_TAKEAWAY           PlaceType = "meal_takeaway"
	K_PLACE_TYPE_MOSQUE                  PlaceType = "mosque"
	K_PLACE_TYPE_MOVIE_RENTAL            PlaceType = "movie_rental"
	K_PLACE_TYPE_MOVIE_THEATER           PlaceType = "movie_theater"
	K_PLACE_TYPE_MOVING_COMPANY          PlaceType = "moving_company"
	K_PLACE_TYPE_MUSEUM                  PlaceType = "museum"
	K_PLACE_TYPE_NIGHT_CLUB              PlaceType = "night_club"
	K_PLACE_TYPE_PAINTER                 PlaceType = "painter"
	K_PLACE_TYPE_PARK                    PlaceType = "park"
	K_PLACE_TYPE_PARKING                 PlaceType = "parking"
	K_PLACE_TYPE_PET_STORE               PlaceType = "pet_store"
	K_PLACE_TYPE_PHARMACY                PlaceType = "pharmacy"
	K_PLACE_TYPE_PHYSIOTHERAPIST         PlaceType = "physiotherapist"
	K_PLACE_TYPE_PLACE_OF_WORSHIP        PlaceType = "place_of_worship"
	K_PLACE_TYPE_PLUMBER                 PlaceType = "plumber"
	K_PLACE_TYPE_POLICE                  PlaceType = "police"
	K_PLACE_TYPE_POST_OFFICE             PlaceType = "post_office"
	K_PLACE_TYPE_REAL_ESTATE_AGENCY      PlaceType = "real_estate_agency"
	K_PLACE_TYPE_RESTAURANT              PlaceType = "restaurant"
	K_PLACE_TYPE_ROOFING_CONTRACTOR      PlaceType = "roofing_contractor"
	K_PLACE_TYPE_RV_PARK                 PlaceType = "rv_park"
	K_PLACE_TYPE_SCHOOL                  PlaceType = "school"
	K_PLACE_TYPE_SHOE_STORE              PlaceType = "shoe_store"
	K_PLACE_TYPE_SHOPPING_MALL           PlaceType = "shopping_mall"
	K_PLACE_TYPE_SPA                     PlaceType = "spa"
	K_PLACE_TYPE_STADIUM                 PlaceType = "stadium"
	K_PLACE_TYPE_STORAGE                 PlaceType = "storage"
	K_PLACE_TYPE_STORE                   PlaceType = "store"
	K_PLACE_TYPE_SUBWAY_STATION          PlaceType = "subway_station"
	K_PLACE_TYPE_SYNAGOGUE               PlaceType = "synagogue"
	K_PLACE_TYPE_TAXI_STAND              PlaceType = "taxi_stand"
	K_PLACE_TYPE_TRAIN_STATION           PlaceType = "train_station"
	K_PLACE_TYPE_TRANSIT_STATION         PlaceType = "transit_station"
	K_PLACE_TYPE_TRAVEL_AGENCY           PlaceType = "travel_agency"
	K_PLACE_TYPE_UNIVERSITY              PlaceType = "university"
	K_PLACE_TYPE_VETERINARY_CARE         PlaceType = "veterinary_care"
	K_PLACE_TYPE_ZOO                     PlaceType = "zoo"
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
	Types     PlaceType
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
		v.Add("types", string(this.Types))
	}
	if len(this.PageToken) > 0 {
		v.Add("pageToken", this.PageToken)
	}
	return v
}

type PlaceTextSearchParam struct {
	Query     string
	Latitude  float64
	Longitude float64
	Radius    int

	Language string
	//MinPrice int
	//MaxPrice int
	OpenNow   bool
	Types     PlaceType
	PageToken string
}

func (this PlaceTextSearchParam) Params() url.Values {
	var v = url.Values{}

	if this.Radius > 0 {
		v.Add("location", fmt.Sprint(this.Latitude, ",", this.Longitude))
		v.Add("radius", fmt.Sprint(this.Radius))
	}
	if len(this.Language) > 0 {
		v.Add("language", this.Language)
	}
	if len(this.Query) > 0 {
		v.Add("query", this.Query)
	}
	if this.OpenNow {
		v.Add("opennow", fmt.Sprint(this.OpenNow))
	}
	if len(this.Types) > 0 {
		v.Add("types", string(this.Types))
	}
	if len(this.PageToken) > 0 {
		v.Add("pageToken", this.PageToken)
	}
	return v
}

type PlaceRadarSearchParam struct {
	Latitude  float64
	Longitude float64
	Radius    int

	Keyword  string
	Language string
	//MinPrice int
	//MaxPrice int
	Name      string
	OpenNow   bool
	Types     PlaceType
	PageToken string
}

func (this PlaceRadarSearchParam) Params() url.Values {
	var v = url.Values{}
	v.Add("location", fmt.Sprint(this.Latitude, ",", this.Longitude))
	if this.Radius > 0 {
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
		v.Add("types", string(this.Types))
	}
	if len(this.PageToken) > 0 {
		v.Add("pageToken", this.PageToken)
	}
	return v
}

type PlaceSearchResults struct {
	HtmlAttributions []string          `json:"html_attributions"`
	Results          []*PlaceBasicInfo `json:"results"`
	NextPageToken    string            `json:"next_page_token"`
	Status           Status            `json:"status"`
}

type PlaceBasicInfo struct {
	Id               string        `json:"id"`
	PlaceId          string        `json:"place_id"`
	Name             string        `json:"name"`
	Geometry         *Geometry     `json:"geometry"`
	Icon             string        `json:"icon"`
	Photos           []*Photo      `json:"photos"`
	Rating           float32       `json:"rating"`
	Reference        string        `json:"reference"`
	Scope            string        `json:"scope"`
	Types            []PlaceType   `json:"types"`
	Vicinity         string        `json:"vicinity"`
	OpeningHours     *OpeningHours `json:"opening_hours"`
	FormattedAddress string        `json:"formatted_address"`
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
	Status           Status        `json:"status"`
	Result           *PlaceDetails `json:"result"`
}

type PlaceDetails struct {
	*PlaceBasicInfo
	AddressComponents        []*AddressComponent `json:"address_components"`
	AdrAddress               string              `json:"adr_address"`
	FormattedPhoneNumber     string              `json:"formatted_phone_number"`
	InternationalPhoneNumber string              `json:"international_phone_number"`
	Reviews                  []*Review           `json:"reviews"`
	URL                      string              `json:"url"`
	UTCOffset                int                 `json:"utc_offset"`
	Website                  string              `json:"website"`
}
