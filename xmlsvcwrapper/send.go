package xmlsvcwrapper

import (
	"bytes"
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

// send function execute the http request with the provided information from Request instance
func (r *Request) send() (*Response, error) {

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

	rawResp, err := io.ReadAll(resp.Body)
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
