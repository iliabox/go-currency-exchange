package config

import "os"

type Config struct {
	DatabaseURL string
	HTTPAddr    string
}

func Load() Config {
	db := os.Getenv("DATABASE_URL")
	if db == "" {
		db = "postgres://symfony:pswd@db:5432/app?sslmode=disable"
	}
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	return Config{DatabaseURL: db, HTTPAddr: addr}
}
