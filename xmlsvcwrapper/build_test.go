package xmlsvcwrapper

import (
	"net/http"
	"testing"
)

func TestRequest_build(t *testing.T) {
	type fields struct {
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

	client := New()

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	testFields := fields{
		Url:         "http://127.0.0.1:3000/test",
		Header:      headers,
		client:      client,
		SoapEnv:     "http://schemas.xmlsoap.org/soap/envelope/",
		SoapType:    "http://www.ibsplc.com/iloyal/member/memberprofiledetail/type/",
		BodyType:    "type:example",
		BodyContent: "Hello world!",
		SecurityHeader: struct {
			UserName   string
			Password   string
			TimeToLive int64
		}{UserName: "example", Password: "example"},
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "test build()",
			fields:  testFields,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Url:            tt.fields.Url,
				Header:         tt.fields.Header,
				client:         tt.fields.client,
				SoapEnv:        tt.fields.SoapEnv,
				SoapType:       tt.fields.SoapType,
				BodyType:       tt.fields.BodyType,
				BodyContent:    tt.fields.BodyContent,
				payloadRequest: tt.fields.payloadRequest,
				SecurityHeader: tt.fields.SecurityHeader,
			}
			if err := r.build(); (err != nil) != tt.wantErr {
				t.Errorf("build() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
