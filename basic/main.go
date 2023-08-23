package main

import (
	"github.com/strernd/tipo/basic/controllers/offer"
	"github.com/strernd/tipo/basic/src/server"
)

func main() {
	oc := offer.NewOfferController()
	server.RunServer(oc)
}
