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
func (p Product) Schema() *nero.Schema {
	return nero.NewSchemaBuilder(&p).
		PkgName("productrepo").
		Collection("products").
		Identity(
			nero.NewColumnBuilder("id", p.ID).
				StructField("ID").Auto().Build(),
		).
		Columns(
			nero.NewColumnBuilder("name", p.Name).Build(),
			nero.NewColumnBuilder("created_at", p.CreatedAt).Auto().Build(),
			nero.NewColumnBuilder("updated_at", p.UpdatedAt).Build(),
		).
		Build()
}
