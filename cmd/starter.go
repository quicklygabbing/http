package cmd

import (
	"github.com/quicklygabbing/http/pkg/http"

	"github.com/pkg/errors"
)

func Starter(address *string) (err error) {
	server := http.NewServer(address)
	err = server.Start(nil)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
