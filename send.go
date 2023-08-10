package soap

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Request         *Request
	RawResponse     *http.Response
	payloadResponse []byte
}

func (r *Response) PayloadResponse() []byte {
	if r.RawResponse == nil {
		return []byte{}
	}
	return r.payloadResponse
}

func (r *Response) StatusCode() int {
	if r.RawResponse == nil {
		return 0
	}
	return r.RawResponse.StatusCode
}

func (r *Request) Send( (*Response, error) {

	//Build the request
	err := r.build()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.Url, bytes.NewReader(r.payloadRequest))
	if err != nil {
		log.Fatalf("failed to create request %s", err)
		return nil, err
	}

	//headers
	req.Header = r.Header
	req.Close = true

	resp, err := r.client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := &Response{
		Request:     r,
		RawResponse: resp,
	}

	rawResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respContent := ContentResponse{}
	if err = xml.Unmarshal(rawResp, &respContent); err != nil {
		return nil, err
	}

	respF := XmlResponse{
		SoapEnv: r.SoapEnv,
		Body: struct {
			Content string `xml:",innerxml"`
		}{Content: respContent.Body.Content},
	}

	if response.payloadResponse, err = xml.Marshal(respF); err != nil {
		return nil, err
	}

	return response, nil
}
