package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyle(t *testing.T) {
	tests := []struct {
		name  string
		input string
		attrs []Attr
		want  string
	}{
		{
			name:  "Empty text",
			input: "",
			attrs: []Attr{AttrFgRed},
			want:  "",
		},
		{
			name:  "No attributes",
			input: "Просто текст",
			attrs: []Attr{},
			want:  "Просто текст",
		},
		{
			name:  "Duplicate attribute",
			input: "Текст",
			attrs: []Attr{AttrBold, AttrBold},
			want:  "\x1b[1mТекст\x1b[0m",
		},
		{
			name:  "Single color",
			input: "Ошибка: файл не найден",
			attrs: []Attr{AttrFgRed},
			want:  "\x1b[31mОшибка: файл не найден\x1b[0m",
		},
		{
			name:  "Single font-style",
			input: "Подчеркнутый",
			attrs: []Attr{AttrUnderline},
			want:  "\x1b[4mПодчеркнутый\x1b[0m",
		},
		{
			name:  "Style combining: color + font-style",
			input: "Важная ошибка",
			attrs: []Attr{AttrFgRed, AttrBold},
			want:  "\x1b[31;1mВажная ошибка\x1b[0m",
		},
		{
			name:  "Style combining: color + font-style + font-decoration",
			input: "Успех!",
			attrs: []Attr{AttrFgGreen, AttrBold, AttrUnderline},
			want:  "\x1b[32;1;4mУспех!\x1b[0m",
		},
		{
			name:  "Style combining: color + bg-Color",
			input: "Информация",
			attrs: []Attr{AttrFgBlack, AttrBgBrightCyan},
			want:  "\x1b[30;106mИнформация\x1b[0m",
		},
		{
			name:  "Multi-line input",
			input: "line\nline\nline\n",
			attrs: []Attr{AttrFgRed},
			want:  "\x1b[31mline\nline\nline\n\x1b[0m",
		},
		{
			name:  "With Emoji",
			input: "Hello 👋 World",
			attrs: []Attr{AttrFgRed},
			want:  "\x1b[31mHello 👋 World\x1b[0m",
		},
		{
			name:  "With CJK sumbols",
			input: "你好世界",
			attrs: []Attr{AttrFgYellow},
			want:  "\x1b[33m你好世界\x1b[0m",
		},
		{
			name:  "Long input",
			input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut bibendum sagittis velit, viverra euismod sapien posuere sit amet. Vivamus euismod sed velit nec elementum. Phasellus a gravida nibh. Phasellus rutrum mollis mi, sed hendrerit dolor semper sit amet. Nam et magna id lacus egestas ultrices. Donec dolor justo, ultrices eu ultrices ut, iaculis sit amet mi. Curabitur vel purus ultrices, porta ex eu, luctus mauris. Nullam et nisi viverra, consequat enim id, rhoncus ligula. Ut hendrerit enim vel turpis lacinia eleifend. Integer facilisis aliquam aliquam. Quisque vitae mi imperdiet, tempor turpis sed, pulvinar orci.",
			attrs: []Attr{AttrBgBlue, AttrFgRed},
			want:  "\x1b[44;31mLorem ipsum dolor sit amet, consectetur adipiscing elit. Ut bibendum sagittis velit, viverra euismod sapien posuere sit amet. Vivamus euismod sed velit nec elementum. Phasellus a gravida nibh. Phasellus rutrum mollis mi, sed hendrerit dolor semper sit amet. Nam et magna id lacus egestas ultrices. Donec dolor justo, ultrices eu ultrices ut, iaculis sit amet mi. Curabitur vel purus ultrices, porta ex eu, luctus mauris. Nullam et nisi viverra, consequat enim id, rhoncus ligula. Ut hendrerit enim vel turpis lacinia eleifend. Integer facilisis aliquam aliquam. Quisque vitae mi imperdiet, tempor turpis sed, pulvinar orci.\x1b[0m",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Style(tc.input, tc.attrs...)
			assert.Equal(t, tc.want, result)
		})
	}
}

func TestColorEnvs(t *testing.T) {
	t.Run("With No Color", func(t *testing.T) {
		NoColor = true
		result := Style("Текст", AttrBgGreen)
		assert.Equal(t, result, "Текст")
		NoColor = false
	})

	t.Run("With No Color", func(t *testing.T) {
		ForceColor = true
		result := Style("Текст", AttrFgRed)
		assert.Equal(t, "\x1b[31mТекст\x1b[0m", result)
		ForceColor = false
	})
}

func TestMakeAttrSeq(t *testing.T) {
	tests := []struct {
		name  string
		input []Attr
		want  string
	}{
		{
			name:  "emty argument",
			input: []Attr{},
			want:  "",
		},
		{
			name:  "classic example",
			input: []Attr{AttrBgBlack, AttrFgBrightGreen},
			want:  "40;92",
		},
		{
			name:  "mix attributes and check sorting",
			input: []Attr{AttrBgBrightCyan, AttrFgBrightRed, AttrUnderline},
			want:  "106;91;4",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := makeAttrSeq(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}
}

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
