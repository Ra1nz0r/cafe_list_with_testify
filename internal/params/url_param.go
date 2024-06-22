package params

import (
	"net/url"
)

func UrlParams(city, count string) (fullURL string) {
	params := url.Values{}
	params.Add("city", city)
	params.Add("count", count)
	baseUrl := "http://0.0.0.0:8080/cafe?"
	return baseUrl + params.Encode()
}
