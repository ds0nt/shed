package leveldb_storage

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDir = os.TempDir() + "/" + "leveldbtest"

func TestMain(m *testing.M) {
	err := os.RemoveAll(testDir)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(testDir, 0755)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := os.RemoveAll(testDir)
		if err != nil {
			panic(err)
		}
	}()

	os.Exit(m.Run())
}

func TestNewLevelDBStorage(t *testing.T) {
	db, err := NewLevelDBStorage(testDir)
	assert.NoError(t, err)

	defer func() {
		err := db.db.Close()
		assert.NoError(t, err)
	}()
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	db, _ := NewLevelDBStorage(testDir)

	err := db.Create(ctx, "testCollection", "testKey1", []byte("testValue1"))
	assert.NoError(t, err)

	defer db.db.Close()
	defer os.Remove(testDir)
}

func TestRead(t *testing.T) {
	ctx := context.Background()
	db, _ := NewLevelDBStorage(testDir)

	_ = db.Create(ctx, "testCollection", "testKey1", []byte("testValue1"))

	value, err := db.Read(ctx, "testCollection", "testKey1")
	assert.Equal(t, []byte("testValue1"), value)
	assert.NoError(t, err)

	defer db.db.Close()
	defer os.Remove(testDir)
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	db, _ := NewLevelDBStorage(testDir)

	_ = db.Create(ctx, "testCollection", "testKey1", []byte("testValue1"))
	_ = db.Update(ctx, "testCollection", "testKey1", []byte("testValue2"))

	value, _ := db.Read(ctx, "testCollection", "testKey1")
	assert.Equal(t, []byte("testValue2"), value)

	defer db.db.Close()
	defer os.Remove(testDir)
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	db, _ := NewLevelDBStorage(testDir)

	_ = db.Create(ctx, "testCollection", "testKey1", []byte("testValue1"))
	_ = db.Delete(ctx, "testCollection", "testKey1")

	value, _ := db.Read(ctx, "testCollection", "testKey1")
	assert.Nil(t, value)

	defer db.db.Close()
	defer os.Remove(testDir)
}

func TestList(t *testing.T) {
	ctx := context.Background()
	db, _ := NewLevelDBStorage(testDir)

	_ = db.Create(ctx, "testCollection", "testKey1", []byte("testValue1"))
	_ = db.Create(ctx, "testCollection", "testKey2", []byte("testValue2"))

	keys, _ := db.List(ctx, "testCollection")
	assert.Equal(t, 2, len(keys))

	defer db.db.Close()
	defer os.Remove(testDir)
}
