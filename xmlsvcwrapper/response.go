package xmlsvcwrapper

import "net/http"

// Response struct holds response values of executed request.
type Response struct {
	Request         *Request
	RawResponse     *http.Response
	payloadResponse []byte
}

// PayloadResponse function returns HTTP response as []byte array for the executed request.
func (r *Response) PayloadResponse() []byte {
	if r.RawResponse == nil {
		return []byte{}
	}
	return r.payloadResponse
}

// PayloadRequest function returns HTTP request as []byte array.
func (r *Response) PayloadRequest() []byte {
	return r.Request.payloadRequest
}

// StatusCode method returns the HTTP status code for the executed request.
//
//	Example: 200
func (r *Response) StatusCode() int {
	if r.RawResponse == nil {
		return 0
	}
	return r.RawResponse.StatusCode
}
