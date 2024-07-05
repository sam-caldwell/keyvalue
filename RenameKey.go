package keyvalue

// RenameKey - Return the maximum width of all values in the current KeyValue struct
//
//	If the KeyValue internal state is nil, return false.
//	If the KeyValue does not exist, return is false.
func (kv *KeyValue[KeyType, ValueType]) RenameKey(currKey, newKey KeyType) bool {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	if kv.data != nil {
		if _, ok := kv.data[currKey]; ok {
			kv.data[newKey] = kv.data[currKey]
			delete(kv.data, currKey)
			return true
		}
	}
	return false
}
