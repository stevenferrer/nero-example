// Code generated by nero, DO NOT EDIT.
package productrepo

import (
	"context"
	"reflect"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/stevenferrer/nero"
	"github.com/stevenferrer/nero-example/model"
	"github.com/stevenferrer/nero/aggregate"
	"github.com/stevenferrer/nero/comparison"
	"github.com/stevenferrer/nero/sort"
)

// Repository is an interface that provides the methods
// for interacting with a Product repository
type Repository interface {
	// BeginTx starts a transaction
	BeginTx(context.Context) (nero.Tx, error)
	// Create creates a Product
	Create(context.Context, *Creator) (id int64, err error)
	// CreateInTx creates a Product in a transaction
	CreateInTx(context.Context, nero.Tx, *Creator) (id int64, err error)
	// CreateMany batch creates Products
	CreateMany(context.Context, ...*Creator) error
	// CreateManyInTx batch creates Products in a transaction
	CreateManyInTx(context.Context, nero.Tx, ...*Creator) error
	// Query queries Products
	Query(context.Context, *Queryer) ([]*model.Product, error)
	// QueryTx queries Products in a transaction
	QueryInTx(context.Context, nero.Tx, *Queryer) ([]*model.Product, error)
	// QueryOne queries a Product
	QueryOne(context.Context, *Queryer) (*model.Product, error)
	// QueryOneTx queries a Product in a transaction
	QueryOneInTx(context.Context, nero.Tx, *Queryer) (*model.Product, error)
	// Update updates a Product or many Products
	Update(context.Context, *Updater) (rowsAffected int64, err error)
	// UpdateTx updates a Product many Products in a transaction
	UpdateInTx(context.Context, nero.Tx, *Updater) (rowsAffected int64, err error)
	// Delete deletes a Product or many Products
	Delete(context.Context, *Deleter) (rowsAffected int64, err error)
	// Delete deletes a Product or many Products in a transaction
	DeleteInTx(context.Context, nero.Tx, *Deleter) (rowsAffected int64, err error)
	// Aggregate performs an aggregate query
	Aggregate(context.Context, *Aggregator) error
	// Aggregate performs an aggregate query in a transaction
	AggregateInTx(context.Context, nero.Tx, *Aggregator) error
}

// Creator is a create builder
type Creator struct {
	name      string
	createdAt string
	updatedAt *string
}

// NewCreator returns a Creator
func NewCreator() *Creator {
	return &Creator{}
}

// Name sets the Name field
func (c *Creator) Name(name string) *Creator {
	c.name = name
	return c
}

// CreatedAt sets the CreatedAt field
func (c *Creator) CreatedAt(createdAt string) *Creator {
	c.createdAt = createdAt
	return c
}

// UpdatedAt sets the UpdatedAt field
func (c *Creator) UpdatedAt(updatedAt *string) *Creator {
	c.updatedAt = updatedAt
	return c
}

// Validate validates the fields
func (c *Creator) Validate() error {
	var err error
	if isZero(c.name) {
		err = multierror.Append(err, nero.NewErrRequiredField("name"))
	}

	if isZero(c.createdAt) {
		err = multierror.Append(err, nero.NewErrRequiredField("created_at"))
	}

	return err
}

// Queryer is a query builder
type Queryer struct {
	limit     uint
	offset    uint
	predFuncs []comparison.PredFunc
	sortFuncs []sort.SortFunc
}

// NewQueryer returns a Queryer
func NewQueryer() *Queryer {
	return &Queryer{}
}

// Where applies predicates
func (q *Queryer) Where(predFuncs ...comparison.PredFunc) *Queryer {
	q.predFuncs = append(q.predFuncs, predFuncs...)
	return q
}

// Sort applies sorting expressions
func (q *Queryer) Sort(sortFuncs ...sort.SortFunc) *Queryer {
	q.sortFuncs = append(q.sortFuncs, sortFuncs...)
	return q
}

// Limit applies limit
func (q *Queryer) Limit(limit uint) *Queryer {
	q.limit = limit
	return q
}

// Offset applies offset
func (q *Queryer) Offset(offset uint) *Queryer {
	q.offset = offset
	return q
}

// Updater is an update builder
type Updater struct {
	name      string
	createdAt string
	updatedAt *string
	predFuncs []comparison.PredFunc
}

// NewUpdater returns an Updater
func NewUpdater() *Updater {
	return &Updater{}
}

// Name sets the Name field
func (c *Updater) Name(name string) *Updater {
	c.name = name
	return c
}

// CreatedAt sets the CreatedAt field
func (c *Updater) CreatedAt(createdAt string) *Updater {
	c.createdAt = createdAt
	return c
}

// UpdatedAt sets the UpdatedAt field
func (c *Updater) UpdatedAt(updatedAt *string) *Updater {
	c.updatedAt = updatedAt
	return c
}

// Where applies predicates
func (u *Updater) Where(predFuncs ...comparison.PredFunc) *Updater {
	u.predFuncs = append(u.predFuncs, predFuncs...)
	return u
}

// Deleter is a delete builder
type Deleter struct {
	predFuncs []comparison.PredFunc
}

// NewDeleter returns a Deleter
func NewDeleter() *Deleter {
	return &Deleter{}
}

// Where applies predicates
func (d *Deleter) Where(predFuncs ...comparison.PredFunc) *Deleter {
	d.predFuncs = append(d.predFuncs, predFuncs...)
	return d
}

// Aggregator is an aggregate query builder
type Aggregator struct {
	v         interface{}
	aggFuncs  []aggregate.AggFunc
	predFuncs []comparison.PredFunc
	sortFuncs []sort.SortFunc
	groupBys  []Field
}

// NewAggregator expects a v and returns an Aggregator
// where 'v' argument must be an array of struct
func NewAggregator(v interface{}) *Aggregator {
	return &Aggregator{v: v}
}

// Aggregate applies aggregate functions
func (a *Aggregator) Aggregate(aggFuncs ...aggregate.AggFunc) *Aggregator {
	a.aggFuncs = append(a.aggFuncs, aggFuncs...)
	return a
}

// Where applies predicates
func (a *Aggregator) Where(predFuncs ...comparison.PredFunc) *Aggregator {
	a.predFuncs = append(a.predFuncs, predFuncs...)
	return a
}

// Sort applies sorting expressions
func (a *Aggregator) Sort(sortFuncs ...sort.SortFunc) *Aggregator {
	a.sortFuncs = append(a.sortFuncs, sortFuncs...)
	return a
}

// Group applies group clauses
func (a *Aggregator) GroupBy(fields ...Field) *Aggregator {
	a.groupBys = append(a.groupBys, fields...)
	return a
}

// rollback performs a rollback
func rollback(tx nero.Tx, err error) error {
	rerr := tx.Rollback()
	if rerr != nil {
		err = errors.Wrapf(err, "rollback error: %v", rerr)
	}
	return err
}

// isZero checks if v is a zero-value
func isZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}
