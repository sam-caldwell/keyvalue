package keyvalue

// Interceptor -
func Interceptor[KeyType comparable, ValueType any](sourceFunc func() (KeyValue[KeyType, ValueType], error)) (string, error) {
	kv, err := sourceFunc()
	return kv.ToString(":", "\n"), err
}
