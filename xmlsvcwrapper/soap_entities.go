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
	Text   string  `xml:",chardata"`
	Header *Header `xml:"wsse:Security,omitempty"`
}

type Header struct {
	XMLName        xml.Name      `xml:"wsse:Security"`
	MustUnderstand string        `xml:"soapenv:mustUnderstand,attr"`
	Wsse           string        `xml:"xmlns:wsse,attr"`
	Timestamp      Timestamp     `xml:"wsu:Timestamp"`
	UsernameToken  UsernameToken `xml:"wsse:UsernameToken"`
}

type UsernameToken struct {
	XMLName  xml.Name `xml:"wsse:UsernameToken"`
	ID       string   `xml:"wsu:Id,attr"`
	Wsu      string   `xml:"xmlns:wsu,attr"`
	Username string   `xml:"wsse:Username"`
	Password Password `xml:"wsse:Password"`
	Nonce    string   `xml:"wsse:Nonce"`
	Created  string   `xml:"wsu:Created"`
}

type Password struct {
	XMLName        xml.Name `xml:"wsse:Password"`
	Type           string   `xml:"Type,attr"`
	PasswordDigest string   `xml:",chardata"`
}

type Timestamp struct {
	XMLName xml.Name `xml:"wsu:Timestamp"`
	ID      string   `xml:"wsu:Id,attr"`
	Wsu     string   `xml:"xmlns:wsu,attr"`
	Created string   `xml:"wsu:Created"`
	Expires string   `xml:"wsu:Expires"`
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
