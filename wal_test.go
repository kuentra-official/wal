package wal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CloseWal(wal *WAL) {
	wal.Close()
	os.RemoveAll(wal.options.DirPath)
}
func TestWalWrite(t *testing.T) {
	dir, _ := os.MkdirTemp("", "test1")
	opts := Options{
		DirPath:           dir,
		DiskFileExtension: ".SDF",
		SegmentSize:       32 * KB,
	}
	wal, err := Open(opts)
	assert.Nil(t, err)
	defer CloseWal(wal)
	pos1, err := wal.Write([]byte("hello1"))
	assert.Nil(t, err)
	assert.NotNil(t, pos1)
	pos2, err := wal.Write([]byte("hello2"))
	assert.Nil(t, err)
	assert.NotNil(t, pos2)
	pos3, err := wal.Write([]byte("hello3"))
	assert.Nil(t, err)
	assert.NotNil(t, pos3)

	val, err := wal.Read(pos1)
	assert.Nil(t, err)
	assert.Equal(t, "hello1", string(val))
	val, err = wal.Read(pos2)
	assert.Nil(t, err)
	assert.Equal(t, "hello2", string(val))
	val, err = wal.Read(pos3)
	assert.Nil(t, err)
	assert.Equal(t, "hello3", string(val))
}
