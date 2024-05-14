package tgholidaybot

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"git.foxminded.ua/foxstudent107051/tgholiday/intrenal/configProvider"
	"git.foxminded.ua/foxstudent107051/tgholiday/intrenal/logger"
)

type SLogger interface {
	SLog() *slog.Logger
}

type SLog struct{}

func (L SLog) SLog() *slog.Logger { return logger.GetLogger() }

type Response struct {
	Name        string `json:"name"`
	NameLocal   string `json:"name_local"`
	Language    string `json:"language"`
	Description string `json:"description"`
	Country     string `json:"country"`
	Location    string `json:"location"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	DateYear    string `json:"date_year"`
	DateMonth   string `json:"date_month"`
	DateDay     string `json:"date_day"`
	WeekDay     string `json:"week_day"`
}

func GetTodayHolidays(country string) []Response {
	url := configProvider.GetQueryUrl(country)

	repo := newFactRepository(url)
	return repo.GetTodayHolidayByCountry()
}

type factRepository struct {
	address string
	client  *http.Client
}

func newFactRepository(addr string) *factRepository {
	return &factRepository{
		address: addr,
		client:  http.DefaultClient,
	}
}

func (r *factRepository) GetTodayHolidayByCountry() []Response {
	var response []Response
	req, err := http.NewRequest(http.MethodGet, r.address, nil)
	if err != nil {
		new(SLog).SLog().Error(err.Error())
	}

	res, err := r.client.Do(req)
	if err != nil {
		new(SLog).SLog().Error(err.Error())
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		new(SLog).SLog().Error(err.Error())
	}

	return response
}
