package params

import (
	"net/url"
	"testing"
)

func TestPack(t *testing.T) {
	tests := []struct {
		uri  string
		data struct {
			Labels     []string `http:"l"`
			MaxResults int      `http:"max"`
			Exact      bool     `http:"x"`
		}
		want string
	}{
		{
			"https://example.com/",
			struct {
				Labels     []string `http:"l"`
				MaxResults int      `http:"max"`
				Exact      bool     `http:"x"`
			}{
				[]string{"golang", "programming"},
				100,
				true,
			},
			"https://example.com/",
		},
	}
	for _, test := range tests {
		uri, err := url.Parse(test.uri)
		if err != nil {
			t.Errorf("parse failed: %v", err)
		}
		got, err := Pack(uri, &test.data)
		if err != nil {
			t.Errorf("pack failed: %v", err)
		}
		if got.String() != test.want {
			t.Errorf("Pack(%v) == %q, want %q", test.data, got, test.want)
		}
	}
}
