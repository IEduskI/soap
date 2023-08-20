package xmlsvcwrapper

import (
	"net/http"
	"testing"
)

func TestRequest_generateSecurityHeader(t *testing.T) {
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

	r := Request{SecurityHeader: struct {
		UserName   string
		Password   string
		TimeToLive int64
	}{UserName: "m-portal", Password: "******", TimeToLive: 5}}

	want := r.generateSecurityHeader()
	tests := []struct {
		name   string
		fields fields
		want   Security
	}{
		{
			name:   "Generate security header",
			fields: fields{SecurityHeader: r.SecurityHeader},
			want:   want,
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
			if got := r.generateSecurityHeader(); got.Header.UsernameToken.Password.PasswordDigest == "" {
				t.Errorf("generateSecurityHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
