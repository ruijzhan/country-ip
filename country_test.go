package country_ip

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountry_Has(t *testing.T) {
	cases := []struct {
		c    string
		ip   string
		want bool
		err  error
	}{
		{
			"CN",
			"114.114.114.114",
			true,
			nil,
		},
		{
			"CN",
			"8.8.8.8",
			false,
			nil,
		},
		{
			"NO_SUCH_COUNTRY",
			"8.8.8.8",
			false,
			ErrCountryNotFound,
		},
		{
			"CN",
			"invalid_ip",
			false,
			errors.New(""),
		},
		{
			"LAN",
			"192.168.1.183",
			true,
			nil,
		},
		{
			"LAN",
			"10.100.1.1",
			true,
			nil,
		},
	}
	for _, cs := range cases {
		get, err := country(cs.c).Has(cs.ip)
		if err != nil && cs.err == nil {
			t.Fatal(err)
		}
		assert.Equal(t, get, cs.want)
	}
}
