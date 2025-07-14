package color

import (
	"maps"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeKey(t *testing.T) {
	tests := []struct {
		name  string
		input []Attr
		want  string
	}{
		{
			name:  "empty argument",
			input: []Attr{},
			want:  "",
		},
		{
			name:  "single attribute",
			input: []Attr{AttrFgRed},
			want:  string([]byte{1, 0, 0, 0, 0, 0, 0, 0}),
		},
		{
			name:  "multiple attributes",
			input: []Attr{AttrFgRed, AttrFgGreen},
			want: string([]byte{
				1, 0, 0, 0, 0, 0, 0, 0,
				2, 0, 0, 0, 0, 0, 0, 0,
			}),
		},
		{
			name:  "larger values",
			input: []Attr{256, 65536},
			want: string([]byte{
				0, 1, 0, 0, 0, 0, 0, 0,
				0, 0, 1, 0, 0, 0, 0, 0,
			}),
		},
		{
			name:  "check sorting",
			input: []Attr{AttrFgGreen, AttrFgRed},
			want: string([]byte{
				1, 0, 0, 0, 0, 0, 0, 0,
				2, 0, 0, 0, 0, 0, 0, 0,
			}),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := makeKey(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}
}

func TestCacheGet(t *testing.T) {
	tests := []struct {
		name     string
		initData map[string]string
		input    []Attr
		expected string
		found    bool
	}{
		{
			name:     "empty cache",
			initData: map[string]string{},
			input:    []Attr{AttrFgGreen},
			expected: "",
			found:    false,
		},
		{
			name: "found in cache",
			initData: map[string]string{
				makeKey([]Attr{AttrFgRed}): makeStartSeq([]Attr{AttrFgRed}),
			},
			input:    []Attr{AttrFgRed},
			expected: makeStartSeq([]Attr{AttrFgRed}),
			found:    true,
		},
		{
			name: "not found in cache",
			initData: map[string]string{
				makeKey([]Attr{AttrFgBlue}): makeStartSeq([]Attr{AttrFgBlue}),
			},
			input:    []Attr{AttrFgGreen},
			expected: "",
			found:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cache := newCache()

			cache.mu.Lock()
			maps.Copy(cache.seqs, tc.initData)
			cache.mu.Unlock()

			result, found := cache.get(tc.input)

			assert.Equal(t, tc.found, found)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestCacheSet(t *testing.T) {
	tests := []struct {
		name  string
		attrs []Attr
		seq   string
		found bool
	}{
		{
			name:  "Zero attrs",
			attrs: []Attr{},
			seq:   "",
			found: true,
		},
		{
			name:  "one attr",
			attrs: []Attr{AttrBgBlue},
			seq:   makeStartSeq([]Attr{AttrBgBlue}),
			found: true,
		},
		{
			name:  "more attrs",
			attrs: []Attr{AttrBgGreen, AttrFgBrightYellow, AttrItalic},
			seq:   makeStartSeq([]Attr{AttrBgGreen, AttrFgBrightYellow, AttrItalic}),
			found: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cache := newCache()

			cache.set(tc.attrs, tc.seq)

			key := makeKey(tc.attrs)
			cache.mu.RLock()
			defer cache.mu.RUnlock()

			seq, found := cache.seqs[key]

			assert.Equal(t, tc.found, found)
			assert.Equal(t, tc.seq, seq)
		})
	}
}
