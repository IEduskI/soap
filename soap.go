package soap

import "net/http"

func New() *Client {
	return NewClient(&http.Client{})
}
