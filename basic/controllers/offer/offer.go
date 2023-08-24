package offer

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/strernd/tipo/basic/src/offer"
)

type BaseOffer struct {
	Products *[]string `json:"products"`
	Price    *uint     `json:"price"`    // monthly price net in euro
	Duration *uint     `json:"duration"` // duration in months
	Email    *string   `json:"email"`    // customer email
}

type ResponseOffer struct {
	Id string `json:"id"`
	BaseOffer
}

type OfferController struct {
	repo *offer.OfferRepository
}

func NewOfferController(repo *offer.OfferRepository) *OfferController {
	return &OfferController{repo: repo}
}

func (c *OfferController) GetOffer(_ context.Context, email string) ([]*ResponseOffer, error) {
	offers, err := c.repo.GetOffer(email)
	if err != nil {
		return nil, err
	}
	res := make([]*ResponseOffer, 0, len(offers))
	for _, o := range offers {
		res = append(res, &ResponseOffer{
			Id: o.Id,
			BaseOffer: BaseOffer{
				Products: nilSplit(o.Products, ","),
				Price:    o.Price,
				Duration: o.Duration,
				Email:    o.Email,
			},
		})
	}
	return res, nil
}

func (c *OfferController) CreateOffer(_ context.Context, o BaseOffer) (*ResponseOffer, error) {
	dboffer := offer.DatabaseOffer{
		Id:       uuid.New().String(),
		Products: nilJoin(o.Products, ","),
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

func nilJoin(strs *[]string, sep string) *string {
	if strs == nil {
		return nil
	}
	joined := strings.Join(*strs, sep)
	return &joined
}

func nilSplit(str *string, del string) *[]string {
	if str == nil {
		return nil
	}
	split := strings.Split(*str, del)
	return &split
}
