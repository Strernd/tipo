package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	offer_controller "github.com/strernd/tipo/basic/controllers/offer"
	offer "github.com/strernd/tipo/basic/src/offer"

	"github.com/strernd/tipo/basic/src/server"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}
	offerRepository := offer.NewOfferRepository(db)
	oc := offer_controller.NewOfferController(offerRepository)
	server.RunServer(oc)
}
