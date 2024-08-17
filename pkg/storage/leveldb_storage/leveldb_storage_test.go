package leveldb_storage

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// clear the leveldb dir
	err := os.RemoveAll("/tmp/testdb")
	if err != nil {
		panic(err)
	}

	err = os.Mkdir("/tmp/testdb", 0755)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestNewLevelDBStorage(t *testing.T) {
	db, err := NewLevelDBStorage("/tmp/testdb")
	assert.NoError(t, err)

	defer func() {
		err := db.db.Close()
		assert.NoError(t, err)

		err = os.Remove("/tmp/testdb")
		assert.NoError(t, err)
	}()
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	db, _ := NewLevelDBStorage("/tmp/testdb")

	err := db.Create(ctx, "testCollection", "testKey1", []byte("testValue1"))
	assert.NoError(t, err)

	defer db.db.Close()
	defer os.Remove("/tmp/testdb")
}

func TestRead(t *testing.T) {
	ctx := context.Background()
	db, _ := NewLevelDBStorage("/tmp/testdb")

	_ = db.Create(ctx, "testCollection", "testKey1", []byte("testValue1"))

	value, err := db.Read(ctx, "testCollection", "testKey1")
	assert.Equal(t, []byte("testValue1"), value)
	assert.NoError(t, err)

	defer db.db.Close()
	defer os.Remove("/tmp/testdb")
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	db, _ := NewLevelDBStorage("/tmp/testdb")

	_ = db.Create(ctx, "testCollection", "testKey1", []byte("testValue1"))
	_ = db.Update(ctx, "testCollection", "testKey1", []byte("testValue2"))

	value, _ := db.Read(ctx, "testCollection", "testKey1")
	assert.Equal(t, []byte("testValue2"), value)

	defer db.db.Close()
	defer os.Remove("/tmp/testdb")
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	db, _ := NewLevelDBStorage("/tmp/testdb")

	_ = db.Create(ctx, "testCollection", "testKey1", []byte("testValue1"))
	_ = db.Delete(ctx, "testCollection", "testKey1")

	value, _ := db.Read(ctx, "testCollection", "testKey1")
	assert.Nil(t, value)

	defer db.db.Close()
	defer os.Remove("/tmp/testdb")
}

func TestList(t *testing.T) {
	ctx := context.Background()
	db, _ := NewLevelDBStorage("/tmp/testdb")

	_ = db.Create(ctx, "testCollection", "testKey1", []byte("testValue1"))
	_ = db.Create(ctx, "testCollection", "testKey2", []byte("testValue2"))

	keys, _ := db.List(ctx, "testCollection")
	assert.Equal(t, 2, len(keys))

	defer db.db.Close()
	defer os.Remove("/tmp/testdb")
}
