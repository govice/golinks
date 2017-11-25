package blockmap

import (
	"errors"
	"os"
	"testing"
)

func TestBlockMap_New(t *testing.T) {
	a := New(os.Getenv("TEST_ROOT"))
	if a == nil {
		t.Error(errors.New("Blockmap: failed to make new blockmap"))
	}
}
