package dict

import "encoding/json"

// Item[K comparable, V any] is a key-value pair.
type Item[K comparable, V any] struct {
	Key   K
	Value V
}

// Dict[K comparable, V any] is a dictionary.
type Dict[K comparable, V any] struct {
	d map[K]V
}

// New[K comparable, V any]() returns a new dictionary.
func New[K comparable, V any]() *Dict[K, V] {
	return &Dict[K, V]{make(map[K]V)}
}

func (d *Dict[K, V]) Clone() *Dict[K, V] {
	d2 := New[K, V]()
	for k, v := range d.d {
		d2.Set(k, v)
	}
	return d2
}

func (d *Dict[K, V]) Set(k K, v V) {
	d.d[k] = v
}

// Has returns true if the dictionary contains the given key.
func (d *Dict[K, V]) Has(k K) bool {
	_, ok := d.d[k]
	return ok
}
func (d *Dict[K, V]) SetItem(item Item[K, V]) {
	d.Set(item.Key, item.Value)
}

// Get returns the value, or the default value, if the key is not found.
func (d *Dict[K, V]) Get(k K, def V) V {
	if v, ok := d.d[k]; ok {
		return v
	}
	return def
}

func (d *Dict[K, V]) Delete(k K) {
	delete(d.d, k)
}

func (d *Dict[K, V]) DeleteItem(item Item[K, V]) {
	d.Delete(item.Key)
}

func (d *Dict[K, V]) Len() int {
	return len(d.d)
}

func (d *Dict[K, V]) Keys() []K {
	keys := make([]K, 0, len(d.d))
	for k := range d.d {
		keys = append(keys, k)
	}
	return keys
}

func (d *Dict[K, V]) Values() []V {
	values := make([]V, 0, len(d.d))
	for _, v := range d.d {
		values = append(values, v)
	}
	return values
}

func (d *Dict[K, V]) Iter() chan Item[K, V] {
	ch := make(chan Item[K, V])
	go func() {
		for k, v := range d.d {
			ch <- Item[K, V]{k, v}
		}
		close(ch)
	}()
	return ch
}
func (d *Dict[K, V]) Items() []Item[K, V] {
	items := make([]Item[K, V], 0, len(d.d))
	for k, v := range d.d {
		items = append(items, Item[K, V]{k, v})
	}
	return items
}

func (d *Dict[K, V]) SetDefault(k K, def V) {
	if _, ok := d.d[k]; !ok {
		d.d[k] = def
	}
}
func (d *Dict[K, V]) SetDefaultF(k K, factory func() V) {
	if _, ok := d.d[k]; !ok {
		d.d[k] = factory()
	}
}

func (d *Dict[K, V]) GetF(k K, factory func() V) V {
	if factory != nil {
		return d.Get(k, factory())
	}
	return d.Get(k, *new(V))
}

func (d *Dict[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.d)
}

func (d *Dict[K, V]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &d.d)
}
