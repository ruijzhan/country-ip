package country_ip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIP_From(t *testing.T) {
	cases := []struct {
		ip   string
		c    string
		want bool
	}{
		{
			"114.114.114.114",
			"CN",
			true,
		},
		{
			"192.168.1.183",
			"LAN",
			true,
		},
		{
			"1.1.1.1",
			"AU",
			true,
		},
	}

	for _, cs := range cases {
		r, err := IP(cs.ip).From(cs.c)
		assert.NoError(t, err)
		assert.Equal(t, r, cs.want)
	}
}
