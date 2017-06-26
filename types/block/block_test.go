package block

import (
	"fmt"
	"testing"
)

func TestBinaryConverters(t *testing.T) {
	blk := New(0, "data", nil)
	bin, _ := blk.marshalBinary()
	var out Block
	_ = out.unmarshalBinary(bin)
	fmt.Println(blk, "\n", out)
	if out.Index != blk.Index || out.Timestamp != blk.Timestamp {
		t.Error("Binary Marshal failed")
	}
}
