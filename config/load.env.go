package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig(filenames ...string) error {
	err := godotenv.Load(filenames...)
	
	if err != nil {
		return err
	}
	
	return nil
}

func GetVar(name string) string {
	return os.Getenv(name)
}