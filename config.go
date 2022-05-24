package main

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type MeituanConfig struct {
	Enable      bool    `toml:"enable"`
	AddressID   int     `toml:"address_id"`
	Poi         int     `toml:"poi"`
	CityID      int     `toml:"city_id"`
	Token       string  `toml:"token"`
	Uid         int64   `toml:"uid"`
	UtmTerm     string  `toml:"utm_term"`
	HomepageLng float64 `toml:"homepage_lng"`
	HomepageLat float64 `toml:"homepage_lat"`
}

type DingDongConfig struct {
	Enable    bool   `toml:"enable"`
	AddressID string `toml:"address_id"`
	Poi       string `toml:"poi"`
	CityID    string `toml:"city_id"`
	Token     string `toml:"token"`
	Uid       string `toml:"uid"`
}

type Config struct {
	Debug      bool            `toml:"debug"`
	LogLevel   string          `toml:"log_level"`
	CronEnable bool            `toml:"cron_enable"`
	CronRule   string          `toml:"cron_rule"`
	Media      string          `toml:"media"`
	SleepMs    int             `toml:"sleep_ms"`
	Meituan    *MeituanConfig  `toml:"meituan"`
	DingDong   *DingDongConfig `toml:"dingdong"`
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
