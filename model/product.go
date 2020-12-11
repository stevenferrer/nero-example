package model

import (
	"time"

	"github.com/sf9v/nero"
)

// Product is a product model
type Product struct {
	ID        int64
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// Schema implements nero.Schemaer
func (p *Product) Schema() *nero.Schema {
	return &nero.Schema{
		Pkg:        "repository",
		Collection: "products",
		Columns: []*nero.Column{
			nero.NewColumn("id", p.ID).StructField("ID").Ident().Auto(),
			nero.NewColumn("name", p.Name),
			nero.NewColumn("created_at", p.CreatedAt).Auto(),
			nero.NewColumn("updated_at", p.UpdatedAt),
		},
	}
}
