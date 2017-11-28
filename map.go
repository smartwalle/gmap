package gmap

import (
	"strings"
	"github.com/smartwalle/nox"
	"net/http"
)

const (
	k_GOOGLE_MAP_API_URL = "https://maps.googleapis.com/maps/api"
)

type GoogleMap struct {
	key         string
	accessToken string
	apiDomain   string
}

func New(key, accessToken string) (client *GoogleMap) {
	client = &GoogleMap{}
	client.key = key
	client.accessToken = accessToken
	client.apiDomain = k_GOOGLE_MAP_API_URL
	return client
}

func (this *GoogleMap) BuildAPI(paths ...string) string {
	var path = this.apiDomain
	for _, p := range paths {
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			if strings.HasSuffix(path, "/") {
				path = path + p
			} else {
				if strings.HasPrefix(p, "/") {
					path = path + p
				} else {
					path = path + "/" + p
				}
			}
		}
	}
	return path
}

func (this *GoogleMap) doRequest(method, url string, param GoogleMapParam, result interface{}) (err error) {
	var v = param.Params()
	if len(this.key) > 0 {
		v.Add("key", this.key)
	}

	req := nox.NewRequest(method, url)
	req.SetParams(v)
	if this.accessToken != "" {
		req.SetHeader("Authorization", "Bearer " + this.accessToken)
	}

	rep := req.Exec()

	switch rep.StatusCode {
	case http.StatusOK:
		if result != nil {
			if err = rep.UnmarshalJSON(result); err != nil {
				return err
			}
		}
	default:
		//var e = &ResponseError{}
		//e.Response = rep.Response
		//if len(rep.MustBytes()) > 0 {
		//	if err = rep.UnmarshalJSON(e); err != nil {
		//		return err
		//	}
		//}
		//return e
	}
	return err
}