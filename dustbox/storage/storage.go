package main

import (
	"errors"
	"os"

	"github.com/adrg/xdg"
)

var data = xdg.DataHome + "/totoro/dustbox/storage/data"

func init_storage() (err error) {
	if err = os.MkdirAll(data, 0600); err != nil {
		return errors.New("cannot create storage folder cause of : " + err.Error())
	}

	return
}
