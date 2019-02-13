package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"
)

const (
	fullSize   = 95
	startIndex = 32
)

var (
	ldTag     = "UNSET"
	ldCommit  = "UNSET"
	ldCompVer = "UNSET"
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
	ver bool
}

type tuple struct {
	s string
	i int
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
	flag.BoolVar(&f.ver, "v", false, "display version information")
	flag.Parse()

	if f.ver {
		fmt.Printf("passgen v.%s\n"+
			"© 2019 Ringo Hoffmann (zekro Development)\n"+
			"Commit Hash: %s\n"+
			"Built with:  %s\n",
			ldTag, ldCommit, strings.Replace(ldCompVer, "_", " ", -1))
		os.Exit(0)
	}

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

func concat(s []string) string {
	b := make([]byte, len(s))
	for i, a := range s {
		b[i] = []byte(a)[0]
	}
	return string(b)
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
	lenFilteredSet := big.NewInt(int64(len(filteredSet)))
	c := make(chan tuple, length)
	res := make([]string, length)

	for i := 0; i < int(length); i++ {
		go func(j int) {
			rnd, err := rand.Int(rand.Reader, lenFilteredSet)
			if err != nil {
				panic(err)
			}
			c <- tuple{filteredSet[rnd.Int64()], j}
		}(i)
	}

	var i uint = 1
	for {
		t := <-c
		res[t.i] = t.s
		if i == length {
			break
		}
		i++
	}

	return concat(res)
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
