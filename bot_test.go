package tgholidaybot

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type myFakeService func(*http.Request) (*http.Response, error)

func (s myFakeService) RoundTrip(req *http.Request) (*http.Response, error) {
	return s(req)
}

func TestGetTodayHolidayByCountry(t *testing.T) {
	client := &http.Client{
		Transport: myFakeService(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(`[
						{"name": "New Year's Day","date": "01/01/2023"},
						{"name": "Christmas Eve","date": "12/24/2023"}
						]`),
				),
			}, nil
		}),
	}

	repo := &factRepository{
		client: client,
	}

	got := repo.GetTodayHolidayByCountry()

	assert.Equal(t, got[0].Name, "New Year's Day")
	assert.Equal(t, got[1].Name, "Christmas Eve")
	assert.Equal(t, got[0].Date, "01/01/2023")
	assert.Equal(t, got[1].Date, "12/24/2023")
}
