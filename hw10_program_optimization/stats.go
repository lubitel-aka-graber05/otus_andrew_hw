package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	"github.com/json-iterator/go"
)

type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {

	return countDomains(r, domain)
}

func countDomains(r io.Reader, domain string) (DomainStat, error) {
	scan := bufio.NewScanner(r)
	result := make(DomainStat)

	for scan.Scan() {
		var user User
		if err := jsoniter.Unmarshal(scan.Bytes(), &user); err != nil {
			return nil, err
		}

		if strings.HasSuffix(user.Email, "."+domain) {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}
	return result, nil
}
