package main

import (
	"context"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"reflect"
	"testing"
)

func TestFindBook(t *testing.T) {
	// create a mocked repository.
	var (
		r     = reltest.New()
		books = []AuthorBook{
			{
				ID:         301,
				Title:      "Go言語によるWebアプリケーション開発",
				Price:      3520,
				AuthorName: "Katherine Cox-Buday",
			},
		}
	)
	r.ExpectFindAll(rel.JoinOn("author", "author.id", "book.author_id"), rel.Eq("author.id", 102)).Result(books)
	r.ExpectFindAll(rel.JoinOn("author", "author.id", "book.author_id"), rel.Eq("author.id", 999)).Result([]AuthorBook{})

	type args struct {
		r  rel.Repository
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    []AuthorBook
		wantErr bool
	}{
		{
			name: "1件検索",
			args: args{
				r:  r,
				id: 102,
			},
			want: []AuthorBook{
				{
					ID:         301,
					Title:      "Go言語によるWebアプリケーション開発",
					Price:      3520,
					AuthorName: "Katherine Cox-Buday",
				},
			},
			wantErr: false,
		},
		{
			name: "存在しないキーを指定",
			args: args{
				r:  r,
				id: 999,
			},
			want: []AuthorBook{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindBook(context.Background(), tt.args.r, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBook() got = %v, want %v", got, tt.want)
			}
		})
	}
}
