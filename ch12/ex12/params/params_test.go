package params

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"testing"
)

func isVisaNumber(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("must be a string: %v", v)
	}
	// http://www.regular-expressions.info/creditcard.html
	re := regexp.MustCompile("^4[0-9]{12}([0-9]{3})?$")
	if !re.MatchString(s) {
		return fmt.Errorf("not a visa number: %q", s)
	}
	return nil
}
func TestPack(t *testing.T) {
	tests := []struct {
		req  *http.Request
		want struct {
			Visa string `http:"visa",validate:"visa"`
		}
	}{
		{
			&http.Request{
				Form: url.Values{
					"visa": []string{"4000000000000000"},
				},
			},
			struct {
				Visa string `http:"visa",validate:"visa"`
			}{
				"4000000000000000",
			},
		},
	}
	for _, test := range tests {
		var got struct {
			Visa string `http:"visa",validate:"visa"`
		}
		err := Unpack(test.req, &got, map[string]Validator{
			"visa": isVisaNumber,
		})
		if err != nil {
			t.Errorf("pack failed: %v", err)
		}
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("Unpack(%v) == %q, want %q", test.req, got, test.want)
		}
	}
}
