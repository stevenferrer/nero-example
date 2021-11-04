package model

import (
	"github.com/stevenferrer/nero"
)

// Product is a product model
type Product struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt *string
}

// Schema implements nero.Schemaer
func (p Product) Schema() *nero.Schema {
	return nero.NewSchemaBuilder(&p).
		PkgName("productrepo").
		Table("products").
		Identity(
			nero.NewFieldBuilder("id", p.ID).
				StructField("ID").Auto().Build(),
		).
		Fields(
			nero.NewFieldBuilder("name", p.Name).Build(),
			nero.NewFieldBuilder("created_at", p.CreatedAt).Build(),
			nero.NewFieldBuilder("updated_at", p.UpdatedAt).
				Optional().Build(),
		).
		Templates(nero.NewSQLiteTemplate()).
		Build()
}
