package main

import (
	"github.com/adrg/xdg"
	"github.com/dgraph-io/badger/v4"
)

var metadata *badger.DB

func init_meta() (err error) {
	metadata, err = badger.Open(badger.DefaultOptions(xdg.DataHome + "/totoro/dustbox/meta"))
	return
}
