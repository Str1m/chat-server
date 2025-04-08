package config

import "github.com/joho/godotenv"

func MustLoad(path string) {
	if err := godotenv.Load(path); err != nil {
		panic("can not load config")
	}
}
