package configProvider

import (
	"fmt"
	"io/ioutil"
	"log/slog"
	"time"

	"git.foxminded.ua/foxstudent107051/tgholiday/intrenal/logger"
	"gopkg.in/yaml.v3"
)

type Config struct {
	HolidayBotApiUrl   string
	HolidayBotApiKey   string
	HolidayBotApiEmail string
}

type SLogger interface {
	SLog() *slog.Logger
}

type SLog struct{}

func (L SLog) SLog() *slog.Logger { return logger.GetLogger() }

func GetConfigs() Config {
	cfg := Config{}
	yamlCfg := readYamlConfigs()
	cfg.HolidayBotApiUrl = yamlCfg["holiday_bot_api_url"]
	cfg.HolidayBotApiKey = yamlCfg["holiday_bot_api_key"]
	cfg.HolidayBotApiEmail = yamlCfg["holiday_bot_api_email"]

	return cfg
}

func GetQueryUrl(country string) string {
	return fmt.Sprintf("%s?api_key=%s&email=%s&country=%s&%s",
		GetConfigs().HolidayBotApiUrl,
		GetConfigs().HolidayBotApiKey,
		GetConfigs().HolidayBotApiEmail,
		country,
		getUrlCurrentDate())
}

func getUrlCurrentDate() string {
	currentTime := time.Now()

	return fmt.Sprintf(
		"year=%v&month=%v&day=%v",
		currentTime.Year(),
		int(currentTime.Month()),
		currentTime.Day(),
	)
}

func readYamlConfigs() map[string]string {
	cfg := make(map[string]string)

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		new(SLog).SLog().Error("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		new(SLog).SLog().Error("Unmarshal: %v", err)
	}

	return cfg
}
