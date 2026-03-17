package config

import (
	"flag"
	"time"
)

type DBConfig struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  time.Duration
}

type Config struct {
	Version   string
	Env       string
	StaticDir string
	Port      int
	DB        DBConfig
}

var version = "0.0.1"

func LoadConfig() *Config {
	var cfg Config

	flag.StringVar(&cfg.Version, "version", version, "Application version")
	flag.StringVar(&cfg.Env, "env", "development", "Enviroment (development, staging, production)")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.IntVar(&cfg.Port, "port", 8080, "Webserver port")

	flag.StringVar(&cfg.DB.DSN, "db-dsn", "", "MySQL DSN")
	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 25, "MySQL max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 25, "MySQL max idle connections")
	flag.DurationVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", 15*time.Minute, "MySQL max connection idle time")

	flag.Parse()

	return &cfg
}
