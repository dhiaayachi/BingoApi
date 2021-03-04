package BingoApi

import "net/http"

type MockClient struct {
	Url  string
	Key  string
	resp *http.Response
	err  error
}

func NewMockCLient(resp *http.Response, err error) *MockClient {
	m := new(MockClient)
	m.resp = resp
	m.err = err
	return m
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	m.Url = req.URL.String()
	m.Key = req.Header.Get("Ocp-Apim-Subscription-Key")
	if m.err != nil {
		return nil, m.err
	}
	return m.resp, nil
}
