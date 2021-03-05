package BingoApi

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBingoApi_NewsSearch(t *testing.T) {

	type fields struct {
		ClientKey  string
		respString string
		statusCode int
		respError  error
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
		{"test basic call with empty json resp", fields{"test key", "{}", 200, nil}, args{"microsoft"}, false},
		{"test basic call with invalid json", fields{"test key", "hello", 200, nil}, args{"microsoft"}, true},
		{"test basic call with 404", fields{"test key", "{}", 404, nil}, args{"microsoft"}, true},
		{"test basic call with error from http client", fields{"test key", "{}", 0, errors.New("tcp error")}, args{"microsoft"}, true},
	}
	for _, tt := range tests {
		// create a new reader with that JSON
		r := ioutil.NopCloser(bytes.NewReader([]byte(tt.fields.respString)))
		m := new(MockClient)

		m.On("Do", mock.Anything).Return(&http.Response{
			StatusCode: tt.fields.statusCode,
			Body:       r,
		}, tt.fields.respError).Once()

		t.Run(tt.name, func(t *testing.T) {
			b := New(tt.fields.ClientKey)
			b.Client = m
			_, err := b.NewsSearch(tt.args.q)

			if err != nil && !tt.wantErr {
				t.Errorf("error %s", err)
				return
			}
			if err == nil && tt.wantErr {
				t.Errorf("error %s", err)
				return
			}
			m.AssertExpectations(t)
		})
	}
}
