package models

import "github.com/uptrace/bun"

type Product struct {
	bun.BaseModel `bun:"table:products"`

	ID       int64 `bun:",pk,autoincrement"`
	Name     string
	Category string
	Quantity int64
	// Kart     *Kart `bun:"rel:has-one,join:id=product_id"`
}
