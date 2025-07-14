package color

import (
	"encoding/binary"
	"errors"
	"sort"
	"sync"
)

var ac = newCache()

type cache struct {
	items map[string]string
	mu    sync.RWMutex
}

func newCache() *cache {
	return &cache{
		items: make(map[string]string, 0),
	}
}

func (c *cache) get(attrs []Attr) (string, bool) {
	key, err := makeKey(attrs)
	if err != nil {
		return "", false
	}

	c.mu.RLock()
	defer c.mu.RUnlock()
	value, found := c.items[key]

	return value, found
}

func (c *cache) set(attrs []Attr, seq string) {
	key, err := makeKey(attrs)

	if err == nil {
		c.mu.Lock()
		defer c.mu.Unlock()
		c.items[key] = seq
	}
}

func makeKey(attrs []Attr) (string, error) {
	if len(attrs) == 0 {
		return "", errors.New("empty attribute")
	}

	intAttrs := make([]int, len(attrs))
	buff := make([]byte, 0, len(attrs))

	for i, atrr := range attrs {
		intAttrs[i] = int(atrr)
	}

	sort.Ints(intAttrs)

	for _, b := range intAttrs {
		buff = binary.LittleEndian.AppendUint64(buff, uint64(b))
	}

	return string(buff), nil
}
