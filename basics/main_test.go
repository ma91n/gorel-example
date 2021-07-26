package main

import (
	"context"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/where"
	"reflect"
	"testing"
)

func TestFindBook(t *testing.T) {
	// create a mocked repository.
	var (
		r    = reltest.New()
		book = Book{
			ID:       301,
			Title:    "Go言語によるWebアプリケーション開発",
			Price:    3520,
			AuthorID: 101,
		}
	)
	r.ExpectFind(where.Eq("id", 301)).Result(book)
	r.ExpectFind(where.Eq("id", 401)).NotFound()

	type args struct {
		r  rel.Repository
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    Book
		wantErr error
	}{
		{
			name: "1件検索",
			args: args{
				r:  r,
				id: 301,
			},
			want: Book{
				ID:       301,
				Title:    "Go言語によるWebアプリケーション開発",
				Price:    3520,
				AuthorID: 101,
			},
			wantErr: nil,
		},
		{
			name: "存在しないキーを指定",
			args: args{
				r:  r,
				id: 401,
			},
			want:    Book{},
			wantErr: NotFoundErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindBook(context.Background(), tt.args.r, tt.args.id)
			if err != tt.wantErr {
				t.Errorf("FindBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBook() got = %v, want %v", got, tt.want)
			}
		})
	}
}
