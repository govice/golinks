package blockmap

import (
	"errors"
	"log"
	"os"
	"testing"
)

func TestBlockMap_New(t *testing.T) {
	a := New(os.Getenv("TEST_ROOT"))
	if a == nil {
		t.Error(errors.New("Blockmap: failed to make new blockmap"))
	}
}

func TestBlockMap_Generate(t *testing.T) {
	b := New(os.Getenv("TEST_ROOT"))
	if err := b.Generate(); err != nil {
		t.Error(err)
	}
	log.Println(b.archive)
}
