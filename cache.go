package color

import (
	"encoding/binary"
	"errors"
	"sync"
)

var ansiCache = newCache()

type cache struct {
	items map[string]string
	mu    sync.RWMutex
}

func newCache() *cache {
	return &cache{
		items: make(map[string]string, 0),
	}
}

func (c *cache) Get(attrs []Attr) (string, bool) {
	key, err := makeKey(attrs)
	if err != nil {
		return "", false
	}

	c.mu.RLock()
	defer c.mu.RUnlock()
	value, found := c.items[key]

	return value, found
}

func (c *cache) Set(attrs []Attr, seq string) {
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

	var intAttrs = make([]uint64, len(attrs))
	var buffer = make([]byte, 0, len(attrs))

	for i, atrr := range attrs {
		intAttrs[i] = uint64(atrr)
	}

	for _, b := range intAttrs {
		buffer = binary.LittleEndian.AppendUint64(buffer, b)
	}

	return string(buffer), nil
}
