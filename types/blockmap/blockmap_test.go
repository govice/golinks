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

func TestEqual(t *testing.T) {
	//Initialize A
	a := New(os.Getenv("TEST_ROOT"))
	if err := a.Generate(); err != nil {
		t.Error(err)
	}
	//Initialize B
	b := New(os.Getenv("TEST_ROOT"))
	if err := b.Generate(); err != nil {
		t.Error(err)
	}

	if !Equal(a, b) {
		t.Error(errors.New("Blockmap: failed to evaluate equal blockmaps"))
	}

	c := &BlockMap{}
	if Equal(a, c) {
		t.Error(errors.New("Blockmap: evaluated equality in unequal blockmaps"))
	}

}
