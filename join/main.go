package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/postgres"
	_ "github.com/lib/pq"
)

type AuthorBook struct {
	ID         int    `db:"id"`
	Title      string `db:"title"`
	Price      int    `db:"price"`
	AuthorName string `db:"name"`
}

func (b AuthorBook) Table() string {
	return "book"
}

func main() {
	adapter, _ := postgres.Open("postgres://postgres:postgres123@localhost/rel_test?sslmode=disable")
	defer adapter.Close()
	repo := rel.New(adapter)

	book, err := FindBook(context.Background(), repo, 102)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(book)
}

func FindBook(ctx context.Context, r rel.Repository, id int) ([]AuthorBook, error) {
	var b []AuthorBook
	if err := r.FindAll(ctx, &b, rel.JoinOn("author", "author.id", "book.author_id"), rel.Eq("author.id", id)); err != nil {
		return nil, err
	}
	return b, nil
}
