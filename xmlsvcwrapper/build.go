package xmlsvcwrapper

import (
	"encoding/xml"
	"fmt"
)

// build function make the request with the correct tags from the information provided in the Request instance
func (r *Request) build() error {
	serviceReq := ServiceRequest{
		SoapEnv: r.SoapEnv,
		Type:    r.SoapType,
		Body: Body{
			Content: RequestBody{
				XMLName: xml.Name{
					Local: r.BodyType,
				},
				RequestBody: r.BodyContent,
			},
		},
	}
	if r.SecurityHeader.UserName != "" && r.SecurityHeader.Password != "" {
		serviceReq.Header = r.generateSecurityHeader()
	}

	respXML, err := xml.Marshal(serviceReq)
	if err != nil {
		return fmt.Errorf("failed to marshal XML: %w", err)
	}
	r.payloadRequest = respXML

	return nil
}
