package country_ip

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"math"
	"net"
	"strconv"
	"strings"

	cidr "github.com/yl2chen/cidranger"
)

var (
	mCidr map[country]cidr.Ranger
	// list of country names
	cList []string
)

func initDB() {
	mCidr = make(map[country]cidr.Ranger)
	cList = make([]string, 0, 100)
	data, err := Asset("apnic.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(bytes.NewReader(data))
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}

		c, n, ok := parseLineV4(string(line))
		if !ok {
			continue
		}
		ranger, ok := mCidr[country(c)]
		if !ok {
			ranger = cidr.NewPCTrieRanger()
			mCidr[country(c)] = ranger
			cList = append(cList, c)
		}
		err = ranger.Insert(cidr.NewBasicRangerEntry(*n))
		if err != nil {
			log.Fatal(err)
		}
	}

	lanRanger := cidr.NewPCTrieRanger()
	mCidr["LAN"] = lanRanger
	for _, n := range lanCidr {
		_, r, _ := net.ParseCIDR(n)
		err := lanRanger.Insert(cidr.NewBasicRangerEntry(*r))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func parseLineV4(line string) (string, *net.IPNet, bool) {
	if !strings.HasPrefix(line, "apnic|") {
		return "", nil, false
	}

	tokens := strings.Split(line, "|")
	if tokens[2] != "ipv4" {
		return "", nil, false
	}

	mask, err := strconv.Atoi(tokens[4])
	if err != nil {
		return "", nil, false
	}

	mask = 32 - int(math.Log(float64(mask))/math.Log(2))
	_, r, err := net.ParseCIDR(tokens[3] + "/" + strconv.Itoa(mask))
	if err != nil {
		return "", nil, false
	}
	return tokens[1], r, true
}
