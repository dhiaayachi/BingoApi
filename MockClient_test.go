package BingoApi

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	a := m.Called(req)
	return a.Get(0).(*http.Response), a.Error(1)
}
