package keyvalue

import (
	"github.com/sam-caldwell/pair"
)

// ToOrderedPairList - Convert KeyValue to OrderedPair
func (kv *KeyValue[KeyType, ValueType]) ToOrderedPairList(sorted bool) *pair.OrderedPair[KeyType, ValueType] {
	var list pair.OrderedPair[KeyType, ValueType]

	if kv.data == nil {
		return &list
	}
	kv.lock.Lock()
	for key, value := range kv.data {
		list.Add(key, value)
	}
	if sorted {
		list.SortByKey()
	}
	kv.lock.Unlock()
	return &list
}
