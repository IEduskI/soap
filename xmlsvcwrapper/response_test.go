package xmlsvcwrapper

import (
	"net/http"
	"reflect"
	"testing"
)

func TestResponse_PayloadResponse(t *testing.T) {
	type fields struct {
		Request         *Request
		RawResponse     *http.Response
		payloadResponse []byte
	}

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()

	response := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" >
					   <soapenv:Header/>
					   <soapenv:Body>
						  <Response>
							 <string>Hello World!</string>
						  </Response>
					   </soapenv:Body>
					</soapenv:Envelope>`

	testRequest := &Request{
		Url:    "http://127.0.0.1:3000/test",
		Header: headers,
		client: client,
	}

	testFields := fields{
		Request:         testRequest,
		RawResponse:     &http.Response{},
		payloadResponse: []byte(response),
	}

	test2Fields := fields{
		Request:         testRequest,
		RawResponse:     nil,
		payloadResponse: []byte{},
	}

	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name:   "Test Retrieve Payload Response",
			fields: testFields,
			want:   []byte(response),
		},
		{
			name:   "Test Retrieve Payload Response empty",
			fields: test2Fields,
			want:   []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Request:         tt.fields.Request,
				RawResponse:     tt.fields.RawResponse,
				payloadResponse: tt.fields.payloadResponse,
			}
			if got := r.PayloadResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PayloadResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_StatusCode(t *testing.T) {
	type fields struct {
		Request         *Request
		RawResponse     *http.Response
		payloadResponse []byte
	}

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()

	textResponse := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" >
					   <soapenv:Header/>
					   <soapenv:Body>
						  <Response>
							 <string>Hello World!</string>
						  </Response>
					   </soapenv:Body>
					</soapenv:Envelope>`

	testRequest := &Request{
		Url:    "http://127.0.0.1:3000/test",
		Header: headers,
		client: client,
	}

	// 1. Test 200
	testFields := fields{
		Request:         testRequest,
		RawResponse:     &http.Response{},
		payloadResponse: []byte(textResponse),
	}
	testFields.RawResponse.StatusCode = http.StatusOK

	// 2. Test 500
	test2Fields := fields{
		Request:         testRequest,
		RawResponse:     &http.Response{},
		payloadResponse: []byte(textResponse),
	}
	test2Fields.RawResponse.StatusCode = http.StatusInternalServerError

	//3. Test nil RawResponse
	test3Fields := fields{
		Request:         testRequest,
		RawResponse:     nil,
		payloadResponse: []byte(textResponse),
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Test Retrieve StatusCode 200",
			fields: testFields,
			want:   http.StatusOK,
		},
		{
			name:   "Test Retrieve StatusCode 500",
			fields: test2Fields,
			want:   http.StatusInternalServerError,
		},
		{
			name:   "Test no response",
			fields: test3Fields,
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Request:         tt.fields.Request,
				RawResponse:     tt.fields.RawResponse,
				payloadResponse: tt.fields.payloadResponse,
			}
			if got := r.StatusCode(); got != tt.want {
				t.Errorf("StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
