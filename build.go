package soap

import (
	"encoding/xml"
	"fmt"
)

type Request struct {
	Url         string
	SoapAction  string
	SoapEnv     string
	SoapType    string
	BodyType    string
	BodyContent string
	Header      struct {
		UserName   string
		Password   string
		TimeToLive int64
	}
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

func (r *Request) SetHeader(usr, pass string, ttl ...int64) *Request {

	r.Header.UserName = usr
	r.Header.Password = pass

	switch {
	case len(ttl) > 0:
		r.Header.TimeToLive = ttl[0]
	default:
		r.Header.TimeToLive = 5
	}
	return r
}

func (s *iflyRequest) build(req *Request) ([]byte, error) {
	serviceReq := ServiceRequest{}

	serviceReq.SoapEnv = req.SoapEnv
	serviceReq.Type = req.SoapType
	serviceReq.Header = generateSecurityHeader(req.Header.UserName, req.Header.Password, req.Header.TimeToLive)
	serviceReq.Body.XMLName.Local = req.BodyType
	serviceReq.Body.RequestBody.RequestBody = req.BodyContent

	respXML, err := xml.Marshal(serviceReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal XML: %w", err)
	}

	return respXML, nil
}
