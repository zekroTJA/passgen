package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

const FULLSETSIZE = 95
const SETSTARTINDEX = 32

func createcharset() string {
	var set = ""
	for i := 0; i < FULLSETSIZE; i++ {
		set += string(SETSTARTINDEX + i)
	}
	return set
}

func getArgVal(args []string, key string) string {
	for i := 0; i < len(args); i++ {
		if args[i] == key && i+1 < len(args) {
			return args[i+1]
		}
	}
	return ""
}

func getArgsArePresent(args []string, keys []string) bool {
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(keys); j++ {
			if args[i] == keys[j] {
				return true
			}
		}
	}
	return false
}

func getRandChar(charset string, pattern *regexp.Regexp) string {
	filteredset := pattern.FindAllString(charset, len(charset))
	rnd := int(rand.Float32() * float32(len(filteredset)))
	return filteredset[rnd]
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	args := os.Args[1:]

	if getArgsArePresent(args, []string{"--help", "?", "/?", "/help", "-h", "-help"}) {
		fmt.Println(
			"passgen v.1.0",
			"\n© 2018 Ringo Hoffmann (zekro Development)",
			"\n\n  -l    Length in chars (defaultly 32)",
			"\n  -rx   Regex of matched chars used to form",
			"\n        password (defaultly '.')",
		)
		os.Exit(0)
	}

	var pwlen = 32
	var regex = regexp.MustCompile(`.`)
	var result = ""

	if pwlenres := getArgVal(args, "-l"); pwlenres != "" {
		var err error
		pwlen, err = strconv.Atoi(getArgVal(args, "-l"))
		if err != nil {
			panic("Can not convert entered length value to int")
		}
	}

	if regexres := getArgVal(args, "-rx"); regexres != "" {
		regex = regexp.MustCompile(regexres)
	}

	fullcharset := createcharset()

	for i := 0; i < pwlen; i++ {
		result += getRandChar(fullcharset, regex)
	}

	fmt.Println(result)
}
