package utils

import "github.com/joho/godotenv"

func LoadEnv(filePath string) error {
	if err := godotenv.Load(filePath); err != nil {
		return err
	}
	return nil
}
