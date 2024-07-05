package keyvalue

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Get - return the value for a given key.
func (kv *KeyValue[KeyType, ValueType]) Get(key KeyType) (value ValueType, err error) {
	raw, ok := kv.FindKey(key)
	if !ok {
		var empty ValueType
		return empty, fmt.Errorf(errors.NotFound)
	}
	return raw, nil
}
