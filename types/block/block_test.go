/*
 *Copyright 2017 Kevin Gentile
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
	"log"
	"testing"
)

func TestEqual(t *testing.T) {

	//Test equal blocks
	blkA := New(0, []byte("data"), nil)
	blkB := blkA
	log.Println("Testing Block Equality")
	if !Equal(blkA, blkB) {
		t.Error("Block equivilents do not match")
	}

	//Test unequal blocks
	log.Println("Testing Block Inequality")
	blkC := New(1, []byte("str"), nil)

	if Equal(blkA, blkC) {
		t.Error("Unequivilent blocks returning as matching")
	}

}
