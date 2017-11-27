package blockmap

import (
	"errors"
	"os"
	"testing"
)

func TestBlockMap_New(t *testing.T) {
	b := New(os.Getenv("TEST_ROOT"))
	if b == nil {
		t.Error(errors.New("Blockmap: failed to make new blockmap"))
	}
}

func TestBlockMap_Generate(t *testing.T) {
	b := New(os.Getenv("TEST_ROOT"))
	if err := b.Generate(); err != nil {
		t.Error(err)
	}
}

func TestBlockMap_PrintBlockMap(t *testing.T) {
	b := New(os.Getenv("TEST_ROOT"))
	if err := b.Generate(); err != nil {
		t.Error(err)
	}
	b.PrintBlockMap()
}
