package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/postgres"
	"github.com/go-rel/rel/where"
	_ "github.com/lib/pq"
	"log"
)

var NotFoundErr = errors.New("not found")

//type Author struct {
//	ID   int    `db:"id"`
//	Name string `db:"name"`
//}

type Book struct {
	ID       int    `db:"id"`
	Title    string `db:"title"`
	Price    int    `db:"price"`
	AuthorID int    `db:"author_id"`
}

func (b Book) Table() string {
	return "book"
}

func main() {
	adapter, err := postgres.Open("postgres://postgres:postgres123@localhost/rel_test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer adapter.Close()
	repo := rel.New(adapter)

	book, err := FindBook(context.Background(), repo, 301)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(book)
}

func FindBook(ctx context.Context, r rel.Repository, id int) (Book, error) {
	var b Book
	if err := r.Find(ctx, &b, where.Eq("id", id)); err != nil {
		if errors.Is(err, rel.NotFoundError{}) {
			return Book{}, NotFoundErr
		}
		return Book{}, err
	}
	return b, nil
}
