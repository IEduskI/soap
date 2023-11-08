package xmlsvcwrapper

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// send function execute the http request with the provided information from Request instance
func (r *Request) send() (*Response, error) {

	req, err := http.NewRequestWithContext(r.Ctx, http.MethodPost, r.Url, bytes.NewReader(r.payloadRequest))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	//headers
	req.Header = r.Header

	/*if r.Ctx != nil {
		req = req.WithContext(r.Ctx)
	} else {
		log.Print("Warning: is higly recommended set the request context")
	}*/

	req.Close = true
	resp, err := r.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request sent: %s, error: %w", string(r.payloadRequest), err)
	}

	defer resp.Body.Close()

	response := &Response{
		Request:     r,
		RawResponse: resp,
	}

	rawResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response, request sent: %s, error: %w", string(r.payloadRequest), err)
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
