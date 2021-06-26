package country_ip

type IP string

// From determins if the IP is in given country's CIDR
func (ip IP) From(c string) (bool, error) {
	return country(c).Has(string(ip))
}
