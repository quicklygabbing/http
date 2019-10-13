package main

import (
	"log"

	"github.com/quicklygabbing/http/cmd"

	"github.com/go-ini/ini"
)

func main() {
	cfg, err := ini.Load(`/quicklygabbing.ini`)
	if err != nil {
		log.Fatalln(err)
	}

	address := cfg.Section(`http`).Key("address").String()
	err = cmd.Starter(&address)
	if err != nil {
		log.Fatal(err)
	}
}
