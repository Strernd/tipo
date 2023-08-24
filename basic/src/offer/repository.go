package offer

import (
	"github.com/jmoiron/sqlx"
)

type DatabaseOffer struct {
	Id       string  `db:"id"`
	Products *string `db:"products"`
	Price    *uint   `db:"price"`    // monthly price net in euro
	Duration *uint   `db:"duration"` // duration in months
	Email    *string `db:"email"`    // customer email
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

func (r *OfferRepository) GetOffer(email string) ([]DatabaseOffer, error) {
	q := "SELECT * FROM offers WHERE email = ?"
	rows, err := r.db.Queryx(q, email)
	if err != nil {
		return nil, err
	}
	var offers []DatabaseOffer
	for rows.Next() {
		var o DatabaseOffer
		err = rows.StructScan(&o)
		if err != nil {
			return nil, err
		}
		offers = append(offers, o)

	}
	return offers, nil
}
