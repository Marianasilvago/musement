package config

import (
	"flag"
	"os"

	"github.com/musement/pkg/envutils"
	"github.com/pkg/errors"
)

type Config struct {
	Environment   string
	WeatherAPIKey string
	Latitude      float64
	Longitude     float64
	ForcastDays   int
}

func GetConfig() (*Config, error) {
	//Default configurations values
	cfg := &Config{
		Environment:   envutils.Development,
		WeatherAPIKey: "796b9a1d3bba40dcb2a94114211802",
		Latitude:      52.374,
		Longitude:     4.9,
		ForcastDays:   2,
	}

	fs := flag.NewFlagSet("Amusement-API", flag.ExitOnError)
	fs.StringVar(&cfg.Environment, "environment", cfg.Environment, "Environment")
	fs.Float64Var(&cfg.Latitude, "latitude", cfg.Latitude, "Latitude")
	fs.Float64Var(&cfg.Longitude, "longitude", cfg.Longitude, "Longitude")
	fs.IntVar(&cfg.ForcastDays, "forecastdays", cfg.ForcastDays, "Forecast Days")
	_ = fs.Parse(os.Args[1:])
	err := checkConfigEnv(cfg)
	if err != nil {
		return nil, errors.WithMessage(err, "check configs")
	}
	return cfg, nil
}

func checkConfigEnv(cfg *Config) error {
	err := envutils.Check(cfg.Environment)
	if err != nil {
		return errors.WithMessage(err, "environment")
	}
	return nil
}
