package offer

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/strernd/tipo/basic/src/offer"
)

type BaseOffer struct {
	Products []string `json:"products"`
	// monthly price net in euro
	Price uint `json:"price"`
	// duration in months
	Duration uint `json:"duration"`
	// customer email
	Email string `json:"email"`
}

type ResponseOffer struct {
	Id string `json:"id"`
	BaseOffer
}

type OfferController struct {
	repo *offer.OfferRepository
}

func NewOfferController(repo *offer.OfferRepository) OfferController {
	return OfferController{repo: repo}
}

func (c OfferController) CreateOffer(_ context.Context, o BaseOffer) (*ResponseOffer, error) {
	dboffer := offer.DatabaseOffer{
		Id:       uuid.New().String(),
		Products: strings.Join(o.Products, ","),
		Price:    o.Price,
		Duration: o.Duration,
		Email:    o.Email,
	}
	err := c.repo.StoreOffer(dboffer)
	if err != nil {
		return nil, err
	}
	res := ResponseOffer{
		BaseOffer: o,
		Id:        dboffer.Id,
	}
	return &res, nil
}
