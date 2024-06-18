package config

import (
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	Port   string `json:"port"`
	Origin string `json:"origin"`
}

var configPath = "/gopx/config.json"

func getConfigPath() string {
	cfgDir := os.Getenv("XDG_CONFIG_HOME")
	if len(cfgDir) == 0 {
		cfgDir = os.Getenv("HOME") + "/.config"
	}

	return cfgDir + configPath
}

func Read() (*Config, error) {
	bs, err := os.ReadFile(getConfigPath())
	if err != nil {
		return nil, err
	}

	var cfg *Config
	if err := json.Unmarshal(bs, &cfg); err != nil {
		return nil, err
	}

	if err := validate(cfg); err != nil {
		return nil, err
	}

	cfg.Port = ":" + cfg.Port

	return cfg, nil
}

func validate(config *Config) error {
	var errs []error
	if _, err := strconv.Atoi(config.Port); err != nil {
		errs = append(errs, errors.New("gopx: provided port value is not a valid Number"))
	}

	if _, err := url.Parse(config.Origin); err != nil {
		errs = append(errs, errors.New("gopx: provided origin value is not a valid URL"))
	}

	return errors.Join(errs...)
}

func Write(config *Config) error {
	if err := validate(config); err != nil {
		return err
	}

	bs, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}

	if err := os.WriteFile(getConfigPath(), bs, 0644); err != nil {
		return err
	}

	return nil
}
