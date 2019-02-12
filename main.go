package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

const (
	fullSize   = 95
	startIndex = 32
)

var strength = []string{
	`[a-z\d]`,
	`[a-zA-Z\d]`,
	`[\w\d]`,
	`[\w\d!"§$%&/()=?-]`,
	`.`,
}

type flags struct {
	len uint
	num uint
	str uint
	rgx string
	sep string
}

func initFlags() *flags {
	f := new(flags)

	strengthStr := ""
	for i, rgx := range strength {
		strengthStr += fmt.Sprintf("  %d - %s\n", i, rgx)
	}

	flag.UintVar(&f.len, "l", 32, "length in characters of the generated token")
	flag.UintVar(&f.num, "n", 1, "number of generated tokens")
	flag.UintVar(&f.str, "s", 2, "strength of the generated token\n"+strengthStr)
	flag.StringVar(&f.rgx, "rx", "", "Regex for used characters for generating the token (matched on the whole charset string)")
	flag.StringVar(&f.sep, "sep", "\n", "custom seperator for multiple tokens")
	flag.Parse()

	if f.len == 0 {
		fmt.Println(`invalid value "0" for flag -l: value must be >= 1`)
		flag.Usage()
	}

	if f.num == 0 {
		fmt.Println(`invalid value "0" for flag -n: value must be >= 1`)
		flag.Usage()
	}

	strengthLen := len(strength)
	if f.str >= uint(strengthLen) {
		fmt.Printf("invalid value \"%d\" for flag -s: value must be in range [0, %d]\n",
			f.str, strengthLen-1)
		flag.Usage()
	}

	return f
}

func getFullCharset() string {
	set := make([]string, fullSize)
	for i := 0; i < fullSize; i++ {
		set[i] = string(startIndex + i)
	}
	return strings.Join(set, "")
}

func getRandString(charset string, length uint, pattern *regexp.Regexp) string {
	filteredSet := pattern.FindAllString(charset, -1)
	res := make([]string, length)
	lenFilteredSet := big.NewInt(int64(len(filteredSet)))

	for i := 0; i < int(length); i++ {
		rnd, err := rand.Int(rand.Reader, lenFilteredSet)
		if err != nil {
			panic(err)
		}
		res[i] = filteredSet[rnd.Int64()]
	}

	return strings.Join(res, "")
}

func generateMultipleTokens(charset string, length uint, pattern *regexp.Regexp, n uint) []string {
	res := make([]string, n)
	for i := 0; i < int(n); i++ {
		res[i] = getRandString(charset, length, pattern)
	}
	return res
}

func main() {
	flags := initFlags()

	fullCharset := getFullCharset()
	pattern := flags.rgx
	if flags.rgx == "" {
		pattern = strength[flags.str]
	}
	rgx := regexp.MustCompile(pattern)

	res := generateMultipleTokens(fullCharset, flags.len, rgx, flags.num)

	fmt.Println(strings.Join(res, flags.sep))
}
