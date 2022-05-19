package models

import "github.com/uptrace/bun"

type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	ID       int64  `bun:",pk,autoincrement"`
	Name     string `bun:"name,notnull"`
	Category string
	Quantity int64
}
