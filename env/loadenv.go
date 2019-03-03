package env

import "github.com/joho/godotenv"

func Loadenv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
