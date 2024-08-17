package leveldb_storage

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/ds0nt/shed/pkg/storage"
	"github.com/syndtr/goleveldb/leveldb"
)

var _ storage.Storer = &LevelDBStorage{}

type LevelDBStorage struct {
	db *leveldb.DB
}

func NewLevelDBStorage(path string) (*LevelDBStorage, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDBStorage{db: db}, nil
}

func (store *LevelDBStorage) Create(ctx context.Context, collection, key string, value []byte) error {
	if store.db == nil {
		return errors.New("database is not initialized")
	}

	err := store.db.Put([]byte(collection+":"+key), value, nil)
	if err != nil {
		return err
	}

	return nil
}
func (store *LevelDBStorage) CreateJSON(ctx context.Context, collection, key string, value interface{}) error {
	if store.db == nil {
		return errors.New("database is not initialized")
	}

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return store.Create(ctx, collection, key, data)
}

func (store *LevelDBStorage) Read(ctx context.Context, collection, key string) ([]byte, error) {
	if store.db == nil {
		return nil, errors.New("database is not initialized")
	}

	value, err := store.db.Get([]byte(collection+":"+key), nil)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (store *LevelDBStorage) ReadJSON(ctx context.Context, collection, key string, value interface{}) error {
	if store.db == nil {
		return errors.New("database is not initialized")
	}

	data, err := store.Read(ctx, collection, key)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, value)
}

func (store *LevelDBStorage) Update(ctx context.Context, collection, key string, value []byte) error {
	if store.db == nil {
		return errors.New("database is not initialized")
	}

	err := store.db.Put([]byte(collection+":"+key), value, nil)
	if err != nil {
		return err
	}

	return nil
}

func (store *LevelDBStorage) Delete(ctx context.Context, collection, key string) error {
	if store.db == nil {
		return errors.New("database is not initialized")
	}

	err := store.db.Delete([]byte(collection+":"+key), nil)
	if err != nil {
		return err
	}

	return nil
}

func (store *LevelDBStorage) List(ctx context.Context, collection string) ([]string, error) {
	if store.db == nil {
		return nil, errors.New("database is not initialized")
	}

	keys := []string{}
	iter := store.db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		if collection == string(key[:len(collection)]) {
			keys = append(keys, string(key))
		}
	}
	iter.Release()

	return keys, nil
}
