package BingoApi

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBingoApi_NewsSearch(t *testing.T) {

	type fields struct {
		ClientKey  string
		respString string
		statusCode int
	}
	type args struct {
		q string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test basic call with empty json resp", fields{"test key", "{}", 200}, args{"microsoft"}, false},
		{"test basic call with invalid json", fields{"test key", "hello", 200}, args{"microsoft"}, true},
		{"test basic call with 404", fields{"test key", "{}", 404}, args{"microsoft"}, true},
	}
	for _, tt := range tests {
		// create a new reader with that JSON
		r := ioutil.NopCloser(bytes.NewReader([]byte(tt.fields.respString)))
		mock := NewMockCLient(&http.Response{
			StatusCode: tt.fields.statusCode,
			Body:       r,
		}, nil)
		t.Run(tt.name, func(t *testing.T) {
			b := &BingoApi{
				ClientKey: tt.fields.ClientKey,
				Client:    mock,
			}
			_, err := b.NewsSearch(tt.args.q)
			if mock.Url != "https://api.bing.microsoft.com/v7.0/news/search?freshness=Day&q="+tt.args.q {
				t.Errorf("Invalid URL %s", mock.Url)
				return
			}
			if mock.Key != "test key" {
				t.Errorf("Invalid KEY %s", mock.Key)
				return
			}
			if err != nil && !tt.wantErr {
				t.Errorf("error %s", err)
				return
			}
		})
	}
}
