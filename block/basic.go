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
	"encoding/json"
)

//Basic is the skeleton used to implement a block
type Basic struct {
	index      int
	timestamp  int64
	data       []byte
	parenthash []byte
	blockhash  []byte
}

// Index returns the index of the block
func (block Basic) Index() int { return block.index }

// Timestamp returns the timestamp of the block
func (block Basic) Timestamp() int64 { return block.timestamp }

// Hash returns the block's hash
func (block Basic) Hash() []byte { return block.blockhash }

// Parenthash returns the parent block's hash
func (block Basic) Parenthash() []byte { return block.parenthash }

// Blockhash returns this blocks hash
func (block Basic) Blockhash() []byte { return block.blockhash }

// Data returns the blocks data
func (block Basic) Data() []byte { return block.data }

// MarshalJSON marshals a block JSON
func (block *Basic) MarshalJSON() ([]byte, error) {
	return json.Marshal(block.JSON())
}

// UnmarshalJSON unmarshals a block JSON
func (block *Basic) UnmarshalJSON(bytes []byte) error {
	holder := &BlockJSON{}

	if err := json.Unmarshal(bytes, &holder); err != nil {
		return err
	}

	block.index = holder.Index
	block.timestamp = holder.Timestamp
	block.parenthash = holder.Parenthash
	block.blockhash = holder.Blockhash
	block.data = holder.Data

	return nil
}

// Serialize returns the json for a block object
func (block Basic) Serialize() ([]byte, error) {
	jsonBlock := BlockJSON{
		Index:      block.Index(),
		Timestamp:  block.Timestamp(),
		Parenthash: append([]byte{}, block.Parenthash()...),
		Data:       append([]byte{}, block.Data()...),
	}
	jsonBytes, err := json.Marshal(jsonBlock)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

// BlockJSON returns the block's json structure
func (block Basic) JSON() BlockJSON {
	return BlockJSON{
		Index:      block.Index(),
		Timestamp:  block.Timestamp(),
		Parenthash: block.Parenthash(),
		Blockhash:  block.Blockhash(),
		Data:       block.Data(),
	}
}

// Block returns the basic block from the JSON structure
func (block BlockJSON) Block() Basic {
	return Basic{
		index:      block.Index,
		timestamp:  block.Timestamp,
		data:       append([]byte{}, block.Data...),
		parenthash: append([]byte{}, block.Parenthash...),
		blockhash:  append([]byte{}, block.Blockhash...),
	}
}

func (block Basic) computeHash() []byte {
	panic("Basic block cannot compute hash")
	return nil
}
