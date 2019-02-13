/*
 *Copyright 2018-2019 Kevin Gentile
 *
 *Licensed under the Apache License, Version 2.0 (the "License");
 *you may not use this file except in compliance with the License.
 *You may obtain a copy of the License at
 *
 *http://www.apache.org/licenses/LICENSE-2.0
 *
 *Unless required by applicable law or agreed to in writing, software
 *distributed under the License is distributed on an "AS IS" BASIS,
 *WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *See the License for the specific language governing permissions and
 *limitations under the License.
 */

package block

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"
)

func TestEqual(t *testing.T) {

	//Test equal blocks
	blkA := NewSHA512(0, []byte("data"), nil)
	blkB := blkA
	log.Println("Testing Block Equality")
	if !bytes.Equal(blkA.Hash(), blkB.Hash()) {
		t.Fail()
	}

	//Test unequal blocks
	log.Println("Testing Block Inequality")
	blkC := NewSHA512(1, []byte("str"), nil)

	if bytes.Equal(blkA.Hash(), blkC.Hash()) {
		t.Error("Unequivilent blocks returning as matching")
	}
}

func TestJSON(t *testing.T) {
	log.Println("Testing block JSON")
	blkA := NewSHA512(0, []byte("GENESIS"), nil)
	blkB := NewSHA512(1, []byte("data"), blkA.Blockhash())

	jsonBytes, err := json.Marshal(blkB)
	if err != nil {
		t.Error(err)
	}

	var blockOut SHA512
	if err := json.Unmarshal(jsonBytes, &blockOut); err != nil {
		t.Error(err)
	}

	if blkB.Index() != blockOut.Index() || blkB.Timestamp() != blockOut.Timestamp() {
		t.Fail()
	}

	if !bytes.Equal(blkB.Parenthash(), blockOut.Parenthash()) {
		t.Fail()
	}

	if !bytes.Equal(blkB.Blockhash(), blockOut.Blockhash()) {
		t.Fail()
	}

	if !bytes.Equal(blkB.Data(), blockOut.Data()) {
		t.Fail()
	}
}

func TestSerialize(t *testing.T) {
	log.Println("Testing block serialize")
	blkA := NewSHA512(0, []byte("GENESIS"), nil)
	blkB := NewSHA512(1, []byte("data"), blkA.Blockhash())

	goldenBytes, err := json.Marshal(blkB)
	if err != nil {
		t.Error(err)
	}

	serializedBytes, err := blkB.Serialize()
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(goldenBytes, serializedBytes) {
		t.Fail()
	}
}
