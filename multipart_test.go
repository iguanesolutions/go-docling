package docling

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"testing"
)

func TestMultipartEncode(t *testing.T) {
	cases := []struct {
		name       string
		s          any
		boundaries uint
		expected   string
	}{
		{
			name: "Bool",
			s: struct {
				Foo bool `json:"foo"`
			}{
				Foo: true,
			},
			boundaries: 1,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"foo\"\r\n\r\ntrue",
		},
		{
			name: "Int",
			s: struct {
				Foo int `json:"foo"`
			}{
				Foo: -1,
			},
			boundaries: 1,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"foo\"\r\n\r\n-1",
		},
		{
			name: "Uint",
			s: struct {
				Foo uint `json:"foo"`
			}{
				Foo: 1,
			},
			boundaries: 1,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"foo\"\r\n\r\n1",
		},
		{
			name: "Float",
			s: struct {
				Foo float64 `json:"foo"`
			}{
				Foo: 1.5,
			},
			boundaries: 1,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"foo\"\r\n\r\n1.5",
		},
		{
			name: "String",
			s: struct {
				Foo string `json:"foo"`
			}{
				Foo: "bar",
			},
			boundaries: 1,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"foo\"\r\n\r\nbar",
		},
		{
			name: "Map",
			s: struct {
				Map map[string]any `json:"map"`
			}{
				Map: map[string]any{
					"foo": "bar",
				},
			},
			boundaries: 1,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"map\"\r\n\r\n{\"foo\":\"bar\"}",
		},
		{
			name: "Struct",
			s: struct {
				Struct struct {
					Foo string `json:"foo"`
				} `json:"struct"`
			}{
				Struct: struct {
					Foo string `json:"foo"`
				}{
					Foo: "bar",
				},
			},
			boundaries: 1,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"struct\"\r\n\r\n{\"foo\":\"bar\"}",
		},
		{
			name: "Array",
			s: struct {
				Slice [2]string `json:"slice"`
			}{
				Slice: [2]string{"foo", "bar"},
			},
			boundaries: 2,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"slice\"\r\n\r\nfoo\r\n--%s\r\nContent-Disposition: form-data; name=\"slice\"\r\n\r\nbar",
		},
		{
			name: "Slice",
			s: struct {
				Slice []string `json:"slice"`
			}{
				Slice: []string{"foo", "bar"},
			},
			boundaries: 2,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"slice\"\r\n\r\nfoo\r\n--%s\r\nContent-Disposition: form-data; name=\"slice\"\r\n\r\nbar",
		},
		{
			name: "Interface",
			s: struct {
				Iface interface{} `json:"iface"`
			}{
				Iface: struct {
					Foo string `json:"foo"`
				}{
					Foo: "bar",
				},
			},
			boundaries: 1,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"iface\"\r\n\r\n{\"foo\":\"bar\"}",
		},
		{
			name: "Pointer",
			s: struct {
				Struct *struct {
					Foo string `json:"foo"`
				} `json:"struct"`
			}{
				Struct: &struct {
					Foo string `json:"foo"`
				}{
					Foo: "bar",
				},
			},
			boundaries: 1,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"struct\"\r\n\r\n{\"foo\":\"bar\"}",
		},
		{
			name: "OmitZero",
			s: struct {
				Foo    string `json:"foo"`
				Struct struct {
					Foo string `json:"foo"`
				} `json:"struct,omitzero"`
			}{
				Foo: "bar",
				Struct: struct {
					Foo string `json:"foo"`
				}{},
			},
			boundaries: 1,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"foo\"\r\n\r\nbar",
		},
		{
			name: "OmitEmpty",
			s: struct {
				Foo string `json:"foo"`
				Bar string `json:"bar,omitempty"`
			}{
				Foo: "bar",
			},
			boundaries: 1,
			expected:   "--%s\r\nContent-Disposition: form-data; name=\"foo\"\r\n\r\nbar",
		},
	}
	for _, c := range cases {
		var b bytes.Buffer
		w := multipart.NewWriter(&b)
		err := multipartEncode(w, c.s)
		if err != nil {
			t.Fatal(err)
		}
		boudaries := make([]any, c.boundaries)
		for i := range c.boundaries {
			boudaries[i] = w.Boundary()
		}
		expected := fmt.Sprintf(c.expected, boudaries...)
		if got := b.String(); got != expected {
			t.Fatalf("\nexpected:\n'''\n%s\n'''\n\ngot:\n'''\n%s\n'''", expected, got)
		}
	}
}
