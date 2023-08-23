package offer

import (
	"context"

	"github.com/google/uuid"
)

type BaseOffer struct {
	Title       string
	Products    []string
	NetPriceEur uint
}

type ResponseOffer struct {
	BaseOffer
	Id string
}

type OfferController struct {
}

func NewOfferController() OfferController {
	return OfferController{}
}

func (c OfferController) CreateOffer(_ context.Context, o BaseOffer) (ResponseOffer, error) {
	// create
	res := ResponseOffer{
		BaseOffer: o,
		Id:        uuid.New().String(),
	}
	return res, nil
}
