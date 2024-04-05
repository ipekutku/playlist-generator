package mocks

import (
	"context"
	"net/http"

	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

type SpotifyServiceMock struct {
	mock.Mock
}

func (m *SpotifyServiceMock) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	ret := m.Called(ctx, code)

	var r0 *oauth2.Token
	if rf, ok := ret.Get(0).(func(context.Context, string) *oauth2.Token); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(*oauth2.Token)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *SpotifyServiceMock) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	ret := m.Called(ctx, token)

	var r0 *http.Client
	if rf, ok := ret.Get(0).(func(context.Context, *oauth2.Token) *http.Client); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(*http.Client)
	}

	return r0
}
