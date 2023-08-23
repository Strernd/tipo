package offer

import (
	"github.com/jmoiron/sqlx"
)

type DatabaseOffer struct {
	Id       string `db:"id"`
	Products string `db:"products"`
	// monthly price net in euro
	Price uint `db:"monthly_price_net_eur"`
	// duration in months
	Duration uint `db:"rent_duration_months"`
	// customer email
	Email string `db:"customer_email"`
}

type OfferRepository struct {
	db *sqlx.DB
}

func NewOfferRepository(db *sqlx.DB) *OfferRepository {
	return &OfferRepository{db: db}
}

func (r *OfferRepository) StoreOffer(offer DatabaseOffer) error {
	q := "INSERT INTO offers (id, products, price, duration, email) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(q, offer.Id, offer.Products, offer.Price, offer.Duration, offer.Email)
	return err
}
