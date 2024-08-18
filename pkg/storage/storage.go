package storage

import "context"

type Storer interface {
	Create(ctx context.Context, collection, key string, value []byte) error
	Read(ctx context.Context, collection, key string) ([]byte, error)
	Update(ctx context.Context, collection, key string, value []byte) error
	List(ctx context.Context, collection string) ([]string, error)
	Delete(ctx context.Context, collection, key string) error

	CreateJSON(ctx context.Context, collection, key string, value interface{}) error
	ReadJSON(ctx context.Context, collection, key string, value interface{}) error
	UpdateJSON(ctx context.Context, collection, key string, value interface{}) error
	ListJSON(ctx context.Context, collection string, values interface{}) error

	Close() error
}
