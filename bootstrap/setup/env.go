package setup

import (
	"fmt"

	"github.com/joho/godotenv"
)

const (
	envfile = ".env"
)

func Env() error {

	var (
		err error
	)

	err = godotenv.Load(envfile)
	if err != nil {
		return fmt.Errorf("error loading .env: %v", err)
	}

	return nil
}
