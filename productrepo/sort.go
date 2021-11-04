// Code generated by nero, DO NOT EDIT.
package productrepo

import (
	"github.com/stevenferrer/nero/sort"
)

// Asc ascending sort direction
func Asc(field Field) sort.SortFunc {
	return func(sorts []*sort.Sort) []*sort.Sort {
		return append(sorts, &sort.Sort{
			Field:     field.String(),
			Direction: sort.Asc,
		})
	}
}

// Desc descending sort direction
func Desc(field Field) sort.SortFunc {
	return func(sorts []*sort.Sort) []*sort.Sort {
		return append(sorts, &sort.Sort{
			Field:     field.String(),
			Direction: sort.Desc,
		})
	}
}
