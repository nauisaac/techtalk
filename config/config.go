package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Marvel Marvel `env-prefix:"MARVEL_"`
}

type Marvel struct {
	BaseURL    string `env:"BASE_URL, required"`
	PublicKey  string `env:"PUBLIC_KEY, required"`
	PrivateKey string `env:"PRIVATE_KEY, required"`
}

var cfg Config

func init() {
	Init("")
}

func Init(path string) {

	godotenv.Load(path)

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal(err)
	}

}

func Set(newConfig Config) Config {
	cfg = newConfig
	return cfg
}

func Get() Config {
	return cfg
}
