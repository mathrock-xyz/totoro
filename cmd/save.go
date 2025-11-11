package main

import (
	"encoding/json"
	"os"

	"github.com/panjf2000/gnet/v2"
	"github.com/spf13/cobra"
)

type inputreq struct {
	Name  string
	Files []string
}

var save = &cobra.Command{
	Use:   "save",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		current, err := os.Getwd()
		if err != nil {
			return
		}

		entries, err := os.ReadDir(current)
		if err != nil {
			return
		}

		isriff := false
		for _, item := range entries {
			if item.IsDir() {
				continue
			}

			if item.Name() == ".riff" {
				isriff = true
				break
			}
		}

		if isriff == false {
		}
	},
}

// input checkpoint name and all the files
func input(req *inputreq) (err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return
	}

	client, err := gnet.NewClient(&gnet.BuiltinEventEngine{})
	if err != nil {
		return
	}
}
