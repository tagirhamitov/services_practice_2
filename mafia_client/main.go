package main

import (
	"flag"
	"log"

	"github.com/tagirhamitov/soa_project/mafia_client/automatic"
	"github.com/tagirhamitov/soa_project/mafia_client/manual"
	"github.com/tagirhamitov/soa_project/mafia_client/parameters"
)

var (
	auto     = flag.Bool("auto", false, "run client in automatic mode")
	username = flag.String("username", "", "username")
	server   = flag.String("server", "", "server address")
)

func main() {
	flag.Parse()
	if *auto {
		params := parameters.ClientParameters{
			Username:      *username,
			ServerAddress: *server,
		}
		err := automatic.Main(params)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		params, err := parameters.GetClientParameters()
		if err != nil {
			log.Fatal(err)
		}
		err = manual.Main(params)
		if err != nil {
			log.Fatal(err)
		}
	}
}
