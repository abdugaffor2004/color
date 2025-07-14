package color

import (
	"encoding/binary"
	"slices"
	"sync"
)

var ac = newCache()

type cache struct {
	seqs map[string]string
	mu   sync.RWMutex
}

func newCache() *cache {
	return &cache{
		seqs: make(map[string]string),
	}
}

func (c *cache) get(attrs []Attr) (string, bool) {
	key := makeKey(attrs)

	c.mu.RLock()
	defer c.mu.RUnlock()
	value, found := c.seqs[key]

	return value, found
}

func (c *cache) set(attrs []Attr, seq string) {
	key := makeKey(attrs)

	c.mu.Lock()
	defer c.mu.Unlock()
	c.seqs[key] = seq
}

func makeKey(attrs []Attr) string {
	if len(attrs) == 0 {
		return ""
	}

	intAttrs := make([]uint64, len(attrs))

	for i, atrr := range attrs {
		intAttrs[i] = uint64(atrr)
	}

	slices.Sort(intAttrs)

	buf := make([]byte, 0, len(attrs))
	for _, b := range intAttrs {
		buf = binary.LittleEndian.AppendUint64(buf, b)
	}

	return string(buf)
}
