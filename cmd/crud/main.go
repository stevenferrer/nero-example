package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/sf9v/nero-example/repository"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = dropTable(db)
		if err != nil {
			log.Fatal(err)
		}
	}()

	repo := repository.NewPostgresRepository(db).Debug()
	ctx := context.Background()
	// create product 1
	product1ID, err := repo.Create(ctx, repository.NewCreator().Name("Product 1"))
	if err != nil {
		log.Fatal(err)
	}

	// create product 2
	_, err = repo.Create(ctx, repository.NewCreator().Name("Product 2"))
	if err != nil {
		log.Fatal(err)
	}

	// query product 1
	product1, err := repo.QueryOne(ctx, repository.NewQueryer().
		Where(repository.IDEq(product1ID)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", product1)

	// update product 1
	now := time.Now()
	_, err = repo.Update(ctx, repository.NewUpdater().
		Name("Updated Product 1").UpdatedAt(&now).
		Where(repository.IDEq(product1ID)))
	if err != nil {
		log.Fatal(err)
	}

	// query product 1 again
	product1, err = repo.QueryOne(ctx, repository.
		NewQueryer().Where(repository.IDEq(product1ID)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", product1)

	// delete product 1
	_, err = repo.Delete(ctx, repository.NewDeleter().
		Where(repository.IDEq(product1ID)))
	if err != nil {
		log.Fatal(err)
	}

	// query remaining products
	products, err := repo.Query(ctx, repository.NewQueryer().Limit(10))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", products[0])
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE products (
		id bigint GENERATED always AS IDENTITY PRIMARY KEY,
		"name" VARCHAR(255) NOT NULL,
		updated_at TIMESTAMP,
		created_at TIMESTAMP DEFAULT now()
	)`)
	return err
}

func dropTable(db *sql.DB) error {
	_, err := db.Exec(`drop table products`)
	return err
}
