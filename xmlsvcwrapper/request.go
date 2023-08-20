package xmlsvcwrapper

import "net/http"

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

// SetUrl function is to set a service url field and its value in the current request.
//
// For Example: To set `http://localhost/lm-ws/lmservices/MemberProfileDetailsService`.
//
//	client.R().
//		SetUrl("http://localhost/lm-ws/lmservices/MemberProfileDetailsService")
func (r *Request) SetUrl(url string) *Request {
	r.Url = url
	return r
}

// SetSoapEnv function is to set a soap envelope field and its value in the current request.
//
// For Example: To set `http://schemas.xmlsoap.org/soap/envelope/`.
//
//	client.R().
//		SetSoapEnv("http://schemas.xmlsoap.org/soap/envelope/")
func (r *Request) SetSoapEnv(env string) *Request {
	r.SoapEnv = env
	return r
}

// SetSoapType function is to set a soap type field and its value in the current request.
//
// For Example: To set `http://www.ibsplc.com/iloyal/member/memberprofiledetail/type/`.
//
//	client.R().
//		SetSoapType("http://www.ibsplc.com/iloyal/member/memberprofiledetail/type/")
func (r *Request) SetSoapType(sType string) *Request {
	r.SoapType = sType
	return r
}

// SetBodyType function is to set a body type field and its value in the current request.
//
// For Example: To set `type:MemberProfileDetailsRequest`.
//
//	client.R().
//		SetBodyType("type:MemberProfileDetailsRequest")
func (r *Request) SetBodyType(bType string) *Request {
	r.BodyType = bType
	return r
}

// SetBodyContent function is to set the content of the request in the current request.
func (r *Request) SetBodyContent(content string) *Request {
	r.BodyContent = content
	return r
}

// SetHeader function is to set a single header field and its value in the current request.
//
// For Example: To set `Content-Type` and `Accept` as `text/xml`.
//
//	client.R().
//		SetHeader("Content-Type", "text/xml; charset=utf-8").
//		SetHeader("Accept", "text/xml; charset=utf-8")
//		SetHeader("SOAPAction", "http://mywebservice.com/km/add/action")
func (r *Request) SetHeader(header, value string) *Request {
	r.Header.Set(header, value)
	return r
}

// SetSecurityHeader function is to set the user, pass and time-to-live for the security header in the current request.
//
// For Example: To set `m-portal` and `******`.
//
//	client.R().
//		SetSecurityHeader("m-portal", "******", 5)
//
// Note: the param ttl is in minutes and is optional if you don't send the default value is '5'
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
