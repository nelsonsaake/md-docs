package bootstrap

import "github.com/nelsonsaake/md-docs/bootstrap/setup"

type BootFunc func() error

var BootRegister = []BootFunc{
	//...
	setup.Env,
}

func Boot() error {

	for _, f := range BootRegister {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}
