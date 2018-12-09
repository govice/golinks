package archivemap

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// ArchiveMap implements marshalling for a well-ordered ordered json map
type ArchiveMap map[string][]byte

// MarshalJSON creates a well ordered JSON byte array for an archive map alphabetically by key
func (am ArchiveMap) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	length := len(am)
	count := 0

	// Sort keys and marshal them to JSON in alphabetical order
	keys := make([]string, 0)
	for k := range am {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		jsonValue, err := json.Marshal(am[key])
		if err != nil {
			return nil, err
		}

		escapedKey := strings.Replace(key, "\\", "/", -1)
		entry := fmt.Sprintf("\"%s\":%s", escapedKey, jsonValue)
		if _, err := buffer.WriteString(entry); err != nil {
			return nil, err
		}
		count++
		if count < length {
			if _, err := buffer.WriteString(","); err != nil {
				return nil, err
			}
		}
	}

	if _, err := buffer.WriteString("}"); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// UnmarshalJSON populates ArchiveMap from a JSON byte array
func (am ArchiveMap) UnmarshalJSON(b []byte) error {
	jsonMap := make(map[string]string)
	err := json.Unmarshal(b, &jsonMap)
	if err != nil {
		return err
	}
	for key, value := range jsonMap {
		fmt.Println("key: " + key + " | value: " + value)
		bytes, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			return err
		}
		am[key] = bytes
	}
	return nil
}
