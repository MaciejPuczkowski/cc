package dict

import "encoding/json"

type GetFSetter[K comparable, V any] interface {
	Set(k K, v V)
	GetF(k K, def func() V) V
}

// Dicter[K comparable, V any] is an interface for a dict.
type Dicter[K comparable, V any] interface {
	GetFSetter[K, V]
	json.Marshaler
	json.Unmarshaler
	SetItem(item Item[K, V])
	SetDefault(k K, def V)
	SetDefaultF(k K, def func() V)
	Get(k K, def V) V
	Delete(k K)
	DeleteItem(item Item[K, V])
	Len() int
	Keys() []K
	Values() []V
	Items() []Item[K, V]
	Clone() *Dict[K, V]
}
