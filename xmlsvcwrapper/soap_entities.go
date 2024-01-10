package xmlsvcwrapper

import "encoding/xml"

// ServiceRequest is the structure for service request
type ServiceRequest struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Text    string   `xml:",chardata"`
	SoapEnv string   `xml:"xmlns:soapenv,attr"`
	Type    string   `xml:"xmlns:type,attr"`
	Header  Security `xml:"soapenv:Header"`
	Body    Body     `xml:"soapenv:Body"`
}

type Body struct {
	Text    string `xml:",chardata"`
	Content RequestBody
}

type RequestBody struct {
	Text        string `xml:",chardata"`
	XMLName     xml.Name
	RequestBody string `xml:",innerxml"`
}

// Security is the structure for the security header
type Security struct {
	Text   string `xml:",chardata,omitempty"`
	Header Header `xml:"wsse:Security,omitempty"`
}

type Header struct {
	XMLName        xml.Name      `xml:"wsse:Security,omitempty"`
	MustUnderstand string        `xml:"soapenv:mustUnderstand,attr,omitempty"`
	Wsse           string        `xml:"xmlns:wsse,attr,omitempty"`
	Timestamp      Timestamp     `xml:"wsu:Timestamp,omitempty"`
	UsernameToken  UsernameToken `xml:"wsse:UsernameToken,omitempty"`
}

type UsernameToken struct {
	XMLName  xml.Name `xml:"wsse:UsernameToken,omitempty"`
	ID       string   `xml:"wsu:Id,attr,omitempty"`
	Wsu      string   `xml:"xmlns:wsu,attr,omitempty"`
	Username string   `xml:"wsse:Username,omitempty"`
	Password Password `xml:"wsse:Password,omitempty"`
	Nonce    string   `xml:"wsse:Nonce,omitempty"`
	Created  string   `xml:"wsu:Created,omitempty"`
}

type Password struct {
	XMLName        xml.Name `xml:"wsse:Password,omitempty"`
	Type           string   `xml:"Type,attr,omitempty"`
	PasswordDigest string   `xml:",chardata,omitempty"`
}

type Timestamp struct {
	XMLName xml.Name `xml:"wsu:Timestamp,omitempty"`
	ID      string   `xml:"wsu:Id,attr,omitempty"`
	Wsu     string   `xml:"xmlns:wsu,attr,omitempty"`
	Created string   `xml:"wsu:Created,omitempty"`
	Expires string   `xml:"wsu:Expires,omitempty"`
}

// ContentResponse saving the response content
type ContentResponse struct {
	Body struct {
		Text    string `xml:",chardata"`
		Content string `xml:",innerxml"`
	} `xml:"Body"`
}

// XmlResponse structure with the correct response format
type XmlResponse struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Text    string   `xml:",chardata"`
	SoapEnv string   `xml:"xmlns:soapenv,attr"`
	Header  string   `xml:"soapenv:Header"`
	Body    struct {
		Content string `xml:",innerxml"`
	} `xml:"soapenv:Body"`
}
