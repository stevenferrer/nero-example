package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/stevenferrer/nero-example/productrepo"
)

func main() {
	dsn := "file:test.db?cache=shared&mode=memory"
	db, err := sql.Open("sqlite3", dsn)
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

	productRepo := productrepo.NewSQLiteRepository(db).Debug()
	ctx := context.Background()
	// create product 1
	now := time.Now().Format(time.RFC3339)
	product1ID, err := productRepo.Create(ctx, productrepo.NewCreator().
		Name("Product 1").CreatedAt(now))
	if err != nil {
		log.Fatal(err)
	}

	// create product 2
	now = time.Now().Format(time.RFC3339)
	_, err = productRepo.Create(ctx, productrepo.NewCreator().
		Name("Product 2").CreatedAt(now))
	if err != nil {
		log.Fatal(err)
	}

	// query product 1
	product1, err := productRepo.QueryOne(ctx, productrepo.NewQueryer().
		Where(productrepo.IDEq(product1ID)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", product1)

	// update product 1
	now = time.Now().Format(time.RFC3339)
	_, err = productRepo.Update(ctx, productrepo.NewUpdater().
		Name("Updated Product 1").UpdatedAt(&now).
		Where(productrepo.IDEq(product1ID)))
	if err != nil {
		log.Fatal(err)
	}

	// query product 1 again
	product1, err = productRepo.QueryOne(ctx, productrepo.
		NewQueryer().Where(productrepo.IDEq(product1ID)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", product1)

	// delete product 1
	_, err = productRepo.Delete(ctx, productrepo.NewDeleter().
		Where(productrepo.IDEq(product1ID)))
	if err != nil {
		log.Fatal(err)
	}

	// query remaining products
	products, err := productRepo.Query(ctx, productrepo.NewQueryer().Limit(10))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", products[0])
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE products (
		id INTEGER PRIMARY KEY,
		"name" TEXT NOT NULL,
		updated_at TEXT,
		created_at TEXT NOT NULL
	)`)
	return err
}

func dropTable(db *sql.DB) error {
	_, err := db.Exec(`drop table products`)
	return err
}
