package models

import "github.com/uptrace/bun"

type Kart struct {
	bun.BaseModel `bun:"table:kart"`

	ID        int64 `bun:",pk,autoincrement"`
	ProductID int64
	Product   *Product `bun:"rel:has-one"`
}
