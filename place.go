package gmap

const (
	k_NEARBY_SEARCH_API = "/place/nearbysearch/json"
)

func (this *GoogleMap) NearbySearch(param NearbySearchParam) (results *NearbySearchResults, err error) {
	var api = this.BuildAPI(k_NEARBY_SEARCH_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}