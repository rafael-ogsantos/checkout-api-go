package models

import "github.com/uptrace/bun"

type Kart struct {
	bun.BaseModel `bun:"table:kart,alias:k"`
	ID            int64 `bun:",pk,autoincrement"`
	ProductId     int64
	Product       []Product `bun:"rel:has-many"`
	Quantity      int64
}
