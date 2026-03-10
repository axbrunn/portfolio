package config

import (
	"flag"
	"time"
)

type Config struct {
	Version   string
	Env       string
	StaticDir string
	Port      int
	db        struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

var version = "0.0.1"

func LoadConfig() *Config {
	var cfg Config

	flag.StringVar(&cfg.Version, "version", version, "Application version")
	flag.StringVar(&cfg.Env, "env", "development", "Enviroment (development, staging, production)")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.IntVar(&cfg.Port, "port", 8080, "Webserver port")

	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "MySQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "MySQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "MySQL max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "MySQL max connection idle time")

	flag.Parse()

	return &cfg
}
