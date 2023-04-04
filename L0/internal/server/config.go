package server

import (
	"encoding/json"
	"os"
)

type config struct {
	Host             string `json:"host"`
	Port             string `json:"port"`
	SubscribeSubject string `json:"subscribe_subject"`
}

type configBox struct {
	Config config `json:"server"`
}

func newConfig(path string) (*config, error) {
	mb, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	conf := configBox{}
	err = json.Unmarshal(mb, &conf)
	if err != nil {
		return nil, err
	}
	return &conf.Config, err
}
