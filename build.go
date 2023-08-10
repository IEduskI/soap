package soap

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Request struct {
	Url            string
	Header         http.Header
	client         *Client
	SoapEnv        string
	SoapType       string
	BodyType       string
	BodyContent    string
	payloadRequest []byte
	SecurityHeader struct {
		UserName   string
		Password   string
		TimeToLive int64
	}
}

func (r *Request) SetUrl(url string) *Request {
	r.Url = url
	return r
}

func (r *Request) SetSoapEnv(env string) *Request {
	r.SoapEnv = env
	return r
}

func (r *Request) SetSoapType(sType string) *Request {
	r.SoapType = sType
	return r
}

func (r *Request) SetBodyType(bType string) *Request {
	r.BodyType = bType
	return r
}

func (r *Request) SetBodyContent(content string) *Request {
	r.BodyContent = content
	return r
}

func (r *Request) SetHeader(header, value string) *Request {
	r.Header.Set(header, value)
	return r
}

func (r *Request) SetSecurityHeader(usr, pass string, ttl ...int64) *Request {

	r.SecurityHeader.UserName = usr
	r.SecurityHeader.Password = pass

	switch {
	case len(ttl) > 0:
		r.SecurityHeader.TimeToLive = ttl[0]
	default:
		r.SecurityHeader.TimeToLive = 5
	}
	return r
}

func (r *Request) build() error {
	serviceReq := ServiceRequest{}

	serviceReq.SoapEnv = r.SoapEnv
	serviceReq.Type = r.SoapType
	serviceReq.Header = r.generateSecurityHeader()
	serviceReq.Body.XMLName.Local = r.BodyType
	serviceReq.Body.RequestBody.RequestBody = r.BodyContent

	respXML, err := xml.Marshal(serviceReq)
	if err != nil {
		return fmt.Errorf("failed to marshal XML: %w", err)
	}
	r.payloadRequest = respXML

	return nil
}
