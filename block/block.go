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
	"github.com/pkg/errors"

	"bytes"
)

// Block the interface used to implement a block
type Block interface {
	computeHash() []byte
	Serialize() ([]byte, error)
	Hash() []byte
	JSON() BlockJSON

	Index() int
	Timestamp() int64
	Parenthash() []byte
	Blockhash() []byte
	Data() []byte
}

type BlockJSON struct {
	Index      int    `json:"index"`
	Timestamp  int64  `json:"timestamp"`
	Data       []byte `json:"data"`
	Parenthash []byte `json:"parentHash"`
	Blockhash  []byte `json:"blockhash,omitempty"`
}

//Validate compares two blocks to verify their parent child relationship.
func Validate(prev, current Block) error {
	if prev.Index()+1 != current.Index() {
		return errors.New("Validate: block indexes do not correlate")
	}
	if !bytes.Equal(prev.Blockhash(), current.Parenthash()) {
		return errors.New("Validate: block hashes do not match")
	}
	h := current.Hash()
	if !bytes.Equal(h, current.Blockhash()) {
		return errors.New("Validate: current block's hash is not valid")
	}
	return nil
}
