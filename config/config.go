package config

import (
	"bufio"
	"os"
	"strings"
)

const (
	filename = ".env"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func MustLoad() *Config {
	env, err := loadEnv(filename)
	if err != nil {
		panic(err)
	}
	return &Config{
		DB: DBConfig{
			DBHost:     env["DB_HOST"],
			DBPort:     env["DB_PORT"],
			DBUser:     env["DB_USER"],
			DBPassword: env["DB_PASSWORD"],
			DBName:     env["DB_NAME"],
			DBSSLMode:  env["DB_SSLMODE"],
		},
	}
}

func loadEnv(filename string) (map[string]string, error) {
	env := map[string]string{}
	fl, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fl.Close()
	scanner := bufio.NewScanner(fl)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		env[key] = value
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return env, nil
}
