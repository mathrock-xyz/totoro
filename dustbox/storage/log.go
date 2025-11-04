package main

import (
	"errors"
	"io"
	"os"

	"github.com/adrg/xdg"
	logger "github.com/charmbracelet/log"
)

var log *logger.Logger // default log

func init_log() (err error) {
	if err = os.MkdirAll(xdg.DataHome+"/totoro/dustbox/storage", 0600); err != nil {
		return errors.New("cannot create storage folder cause of : " + err.Error())
	}

	file, err := os.OpenFile(xdg.DataHome+"/totoro/dustbox/storage/logfile", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("cannot create storage folder cause of : " + err.Error())
	}

	writers := io.MultiWriter(file, os.Stdout)
	log = logger.NewWithOptions(writers, logger.Options{
		Prefix: "Storage ",
	})

	return nil
}
