package gmap

const (
	k_PLACE_NEARBY_SEARCH_API = "/place/nearbysearch/json"
	k_PLACE_TEXT_SEARCH_API = "/place/textsearch/json"
	k_PLACE_DETAILS_API       = "/place/details/json"
)

func (this *GoogleMap) PlaceNearbySearch(param PlaceNearbySearchParam) (results *PlaceSearchResults, err error) {
	var api = this.BuildAPI(k_PLACE_NEARBY_SEARCH_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}

func (this *GoogleMap) PlaceTextSearch(param PlaceTextSearchParam) (results *PlaceSearchResults, err error) {
	var api = this.BuildAPI(k_PLACE_TEXT_SEARCH_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}

func (this *GoogleMap) PlaceDetails(param PlaceDetailsParam) (results *PlaceDetailsResults, err error) {
	var api = this.BuildAPI(k_PLACE_DETAILS_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}