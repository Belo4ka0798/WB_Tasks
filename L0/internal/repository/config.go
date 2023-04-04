package repository

import (
	"encoding/json"
	"os"
)

type config struct {
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type configBox struct {
	Config config `json:"database"`
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
