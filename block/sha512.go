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
	"crypto/sha512"
	"time"
)

// SHA512 is a block that uses the SHA512 hash
type SHA512 struct {
	Basic
}

// NewSHA512 creates a new block using SHA512 hashing
func NewSHA512(index int, data []byte, parentHash []byte) *SHA512 {
	blk := &SHA512{Basic{
		index:      index,
		timestamp:  time.Now().UnixNano(),
		data:       append([]byte{}, data...),
		parenthash: append([]byte{}, parentHash...),
	}}
	blk.computeHash()
	return blk
}

// Hash calculates and assigns the hash to a block.
func (block *SHA512) computeHash() []byte {
	blkhash := sha512.New()
	blockBytes, err := block.Serialize()
	if err != nil {
		panic(err)
	}
	if _, err := blkhash.Write(blockBytes); err != nil {
		panic(err)
	}
	block.blockhash = append([]byte{}, blkhash.Sum(nil)...)
	return block.blockhash
}

// NewGenesis returns a new gensis block hashed with SHA512
func NewSHA512Genesis() *SHA512 {
	genesis := &SHA512{Basic{
		index:      0,
		timestamp:  time.Time{}.UnixNano(),
		data:       append([]byte{}, []byte("GENSIS BLOCK")...),
		parenthash: append([]byte{}, []byte("")...),
	}}
	genesis.computeHash()
	return genesis
}
