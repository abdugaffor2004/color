package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeKey(t *testing.T) {
	tests := []struct {
		name    string
		input   []Attr
		want    string
		wantErr bool
	}{
		{
			name:    "empty argument",
			input:   []Attr{},
			want:    "",
			wantErr: true,
		},
		{
			name:    "single attribute",
			input:   []Attr{AttrFgRed},
			want:    string([]byte{1, 0, 0, 0, 0, 0, 0, 0}),
			wantErr: false,
		},
		{
			name:    "multiple attributes",
			input:   []Attr{AttrFgRed, AttrFgGreen},
			want:    string([]byte{1, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0}),
			wantErr: false,
		},
		{
			name:    "larger values",
			input:   []Attr{256, 65536},
			want:    string([]byte{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}),
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := makeKey(tc.input)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tc.want, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, result)
			}
		})
	}
}

func TestCache(t *testing.T) {

	tests := []struct {
		name    string
		attrs   []Attr
		seq     string
		isFound bool
	}{
		{
			name:    "Zero attrs",
			attrs:   []Attr{},
			seq:     "",
			isFound: false,
		},
		{
			name:    "one attr",
			attrs:   []Attr{AttrBgBlue},
			seq:     makeStartSeq([]Attr{AttrBgBlue}),
			isFound: true,
		},
		{
			name:    "more attrs",
			attrs:   []Attr{AttrBgGreen, AttrFgBrightYellow, AttrItalic},
			seq:     makeStartSeq([]Attr{AttrBgGreen, AttrFgBrightYellow, AttrItalic}),
			isFound: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cache := newCache()

			cache.set(tc.attrs, tc.seq)
			seq, found := cache.get(tc.attrs)

			assert.Equal(t, tc.isFound, found)
			assert.Equal(t, tc.seq, seq)
		})
	}

}
