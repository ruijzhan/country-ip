package country_ip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseLineV4(t *testing.T) {
	cases := []struct {
		line    string
		country string
		cidr    string
		valid   bool
	}{
		{"apnic|CN|ipv4|43.254.228.0|1024|20140729|allocated", "CN", "43.254.228.0/22", true},
		{"apnic|TW|asn|7532|8|19970322|allocated", "", "", false},
		{"apnic|AU|ipv6|2401:700::|32|20110606|allocated", "", "", false},
		{"# statement of the location in which any specific resource may", "", "", false},
	}

	for _, cs := range cases {
		country, cidr, valid := parseLineV4(cs.line)
		assert.Equal(t, cs.country, country)
		assert.Equal(t, cs.cidr, func() string {
			if cidr == nil {
				return ""
			} else {
				return cidr.String()
			}
		}())
		assert.Equal(t, cs.valid, valid)
	}
}
