package interfaces

import "net/http"

type HttpClient interface {
	Get(relativeUrl string) (*http.Response, error)
	Post(relativeUrl string, body interface{}) (*http.Response, error)
}
