package xmlsvcwrapper

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestRequest_SetBodyContent(t *testing.T) {
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
	type args struct {
		content string
	}

	client := New()

	testFields := fields{
		Url:         "http://127.0.0.1:3000/test",
		client:      client,
		BodyContent: "Hello world!",
	}

	testArgs := args{
		content: "Hello world!",
	}

	testRequest := &Request{
		Url:         "http://127.0.0.1:3000/test",
		client:      client,
		BodyContent: "Hello world!",
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test Set Body Content",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
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
			if got := r.SetBodyContent(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBodyContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetBodyType(t *testing.T) {
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
	type args struct {
		bType string
	}

	client := New()

	testFields := fields{
		Url:      "http://127.0.0.1:3000/test",
		client:   client,
		BodyType: "type:example",
	}

	testArgs := args{
		bType: "type:example",
	}

	testRequest := &Request{
		Url:      "http://127.0.0.1:3000/test",
		client:   client,
		BodyType: "type:example",
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test Set Body Type",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
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
			if got := r.SetBodyType(tt.args.bType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBodyType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetHeader(t *testing.T) {
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

	type args struct {
		header string
		value  string
	}

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()

	testFields := fields{
		Url:    "http://127.0.0.1:3000/test",
		Header: headers,
		client: client,
	}

	testArgs := args{
		header: "Content-Type",
		value:  "text/xml; charset=utf-8",
	}

	testRequest := &Request{
		Url:    "http://127.0.0.1:3000/test",
		Header: headers,
		client: client,
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test Set Header",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
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
			if got := r.SetHeader(tt.args.header, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetSecurityHeader(t *testing.T) {
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
	type args struct {
		usr  string
		pass string
		ttl  []int64
	}

	client := New()

	testFields := fields{
		Url:    "http://127.0.0.1:3000/test",
		client: client,
		SecurityHeader: struct {
			UserName   string
			Password   string
			TimeToLive int64
		}{UserName: "example", Password: "example"},
	}

	testArgs := args{
		usr:  "example",
		pass: "example",
	}

	test2Args := args{
		usr:  "example",
		pass: "example",
		ttl:  []int64{5},
	}

	testRequest := &Request{
		Url:    "http://127.0.0.1:3000/test",
		client: client,
		SecurityHeader: struct {
			UserName   string
			Password   string
			TimeToLive int64
		}{UserName: "example", Password: "example", TimeToLive: 5},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test Set Security Header",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
		},
		{
			name:   "Test Set Security Header - set ttl",
			fields: testFields,
			args:   test2Args,
			want:   testRequest,
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
			if got := r.SetSecurityHeader(tt.args.usr, tt.args.pass, tt.args.ttl...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSecurityHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetSoapEnv(t *testing.T) {
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
	type args struct {
		env string
	}

	client := New()

	testFields := fields{
		Url:     "http://127.0.0.1:3000/test",
		client:  client,
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
	}

	testArgs := args{
		env: "http://schemas.xmlsoap.org/soap/envelope/",
	}

	testRequest := &Request{
		Url:     "http://127.0.0.1:3000/test",
		client:  client,
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test Set Soap Envelope",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
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
			if got := r.SetSoapEnv(tt.args.env); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSoapEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetSoapType(t *testing.T) {
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
	type args struct {
		sType string
	}

	client := New()

	testFields := fields{
		Url:      "http://127.0.0.1:3000/test",
		client:   client,
		SoapType: "http://www.ibsplc.com/iloyal/member/memberprofiledetail/type/",
	}

	testArgs := args{
		sType: "http://www.ibsplc.com/iloyal/member/memberprofiledetail/type/",
	}

	testRequest := &Request{
		Url:      "http://127.0.0.1:3000/test",
		client:   client,
		SoapType: "http://www.ibsplc.com/iloyal/member/memberprofiledetail/type/",
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test Soap Type",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
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
			if got := r.SetSoapType(tt.args.sType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSoapType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetUrl(t *testing.T) {
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
	type args struct {
		url string
	}

	client := New()

	testFields := fields{
		Url:    "http://127.0.0.1:3000/test",
		client: client,
	}

	testArgs := args{
		url: "http://127.0.0.1:3000/test",
	}

	testRequest := &Request{
		Url:    "http://127.0.0.1:3000/test",
		client: client,
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test Set Url",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
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
			if got := r.SetUrl(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetContext(t *testing.T) {
	type fields struct {
		Url            string
		Header         http.Header
		client         *Client
		Ctx            context.Context
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
	type args struct {
		ctx context.Context
	}

	client := New()
	ctx := context.Background()

	testFields := fields{
		Url:    "http://127.0.0.1:3000/test",
		client: client,
		Ctx:    ctx,
	}

	testRequest := &Request{
		Url:    "http://127.0.0.1:3000/test",
		client: client,
		Ctx:    ctx,
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test SetContext()",
			fields: testFields,
			args:   args{ctx: ctx},
			want:   testRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Url:            tt.fields.Url,
				Header:         tt.fields.Header,
				client:         tt.fields.client,
				Ctx:            tt.fields.Ctx,
				SoapEnv:        tt.fields.SoapEnv,
				SoapType:       tt.fields.SoapType,
				BodyType:       tt.fields.BodyType,
				BodyContent:    tt.fields.BodyContent,
				payloadRequest: tt.fields.payloadRequest,
				SecurityHeader: tt.fields.SecurityHeader,
			}
			if got := r.SetContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
