package country_ip

import (
	"errors"
	"net"
)

func init() {
	initDB()
}

var (
	// CN represents China
	CN = NewCountry("CN")

	// ErrCountryNotFound determins the given country dose not exit in list
	// https://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest
	ErrCountryNotFound = errors.New("country not found")
)

type Country interface {
	// Has checks if country's CIDR contains given ip
	Has(string) (bool, error)

	// String returns country's name
	String() string
}

func NewCountry(c string) Country {
	return country(c)
}

type country string

func (c country) Has(ip string) (bool, error) {
	r, ok := mCidr[c]
	if !ok {
		return false, ErrCountryNotFound
	}
	i := net.ParseIP(ip)
	return r.Contains(i)
}

func (c country) String() string {
	return string(c)
}

// List returns countries names
func List() []string {
	return cList
}
