package xmlsvcwrapper

import (
	"encoding/xml"
	"fmt"
)

// build function make the request with the correct tags from the information provided in the Request instance
func (r *Request) build() error {
	serviceReq := ServiceRequest{}

	serviceReq.SoapEnv = r.SoapEnv
	serviceReq.Type = r.SoapType
	serviceReq.Header = r.generateSecurityHeader()
	serviceReq.Body.Content.XMLName.Local = r.BodyType
	serviceReq.Body.Content.RequestBody = r.BodyContent

	respXML, err := xml.Marshal(serviceReq)
	if err != nil {
		return fmt.Errorf("failed to marshal XML: %w", err)
	}
	r.payloadRequest = respXML

	return nil
}
