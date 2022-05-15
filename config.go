package main

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Config struct {
	LogLevel    string  `toml:"log_level"`
	AddressID   int     `toml:"address_id"`
	UtmTerm     string  `toml:"utm_term"`
	Poi         int     `toml:"poi"`
	HomepageLng float64 `toml:"homepage_lng"`
	HomepageLat float64 `toml:"homepage_lat"`
	CityID      int     `toml:"city_id"`
	Token       string  `toml:"token"`
	CronEnable  bool    `toml:"cron_enable"`
	CronRule    string  `toml:"cron_rule"`
	Media       string  `toml:"media"`
	SleepMs     int     `toml:"sleep_ms"`
	Uid         int64   `toml:"uid"`
}

func ParseConf(reader io.Reader) (*Config, error) {
	confData, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	confData = bytes.TrimPrefix(confData, []byte{239, 187, 191})
	conf := &Config{}
	err = toml.Unmarshal(confData, conf)
	if err != nil {
		log.Errorf("解析配置文件遇到错误:%s", err.Error())
		return nil, err
	}
	return conf, nil
}
