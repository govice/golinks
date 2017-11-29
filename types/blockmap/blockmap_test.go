package blockmap

import (
	"errors"
	"os"
	"testing"
)

func TestBlockMap_New(t *testing.T) {
	b := New(os.Getenv("TEST_ROOT"))
	if b == nil {
		t.Error(errors.New("blockmap: failed to make new blockmap"))
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
	//b.PrintBlockMap()
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
		t.Error(errors.New("blockmap: failed to evaluate equal blockmaps"))
	}

	c := &BlockMap{}
	if Equal(a, c) {
		t.Error(errors.New("blockmap: evaluated equality in unequal blockmaps"))
	}

}

func TestBlockMap_IO(t *testing.T) {
	//Generate initial blockmap from the test root
	b := New(os.Getenv("TEST_ROOT"))
	if err := b.Generate(); err != nil {
		t.Error(err)
	}

	//Save the blockmap
	if err := b.Save(b.Root); err != nil {
		t.Error(err)
	}
	//Load the blockmap in a new structure
	a := &BlockMap{}
	if err := a.Load(b.Root); err != nil {
		t.Error(err)
	}
	//Ensure both maps are equal
	if !Equal(b, a) {
		t.Error(errors.New("BlockMapIO failed to reload map"))
	}

}
