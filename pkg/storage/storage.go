package storage

import "context"

type Storer interface {
	Create(ctx context.Context, collection, key string, value []byte) error
	CreateJSON(ctx context.Context, collection, key string, value interface{}) error
	Read(ctx context.Context, collection, key string) ([]byte, error)
	ReadJSON(ctx context.Context, collection, key string, value interface{}) error
	Update(ctx context.Context, collection, key string, value []byte) error
	Delete(ctx context.Context, collection, key string) error
	List(ctx context.Context, collection string) ([]string, error)
}
