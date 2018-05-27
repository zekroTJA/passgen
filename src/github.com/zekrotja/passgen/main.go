package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	FULLSETSIZE   = 95
	SETSTARTINDEX = 32
	STRENGTH_0    = `[a-z\d]`
	STRENGTH_1    = `[a-zA-Z\d]`
	STRENGTH_2    = `[\w\d]`
	STRENGTH_3    = `[\w\d!"§$%&/()=?-]`
	STRENGTH_4    = `.`
)

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

func genPasswd(len int, charset string, pattern *regexp.Regexp) string {
	var res = ""
	for i := 0; i < len; i++ {
		res += getRandChar(charset, pattern)
	}
	return res
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	args := os.Args[1:]

	if getArgsArePresent(args, []string{"--help", "?", "/?", "/help", "-h", "-help"}) {
		fmt.Println(
			"passgen v.1.0",
			"\n© 2018 Ringo Hoffmann (zekro Development)",
			"\n\n  -l    Length in chars (defaultly 32)",
			"\n  -n    Number of generated password strings",
			"\n        (defaultly 1)",
			"\n  -s    Strength as number instead of regex",
			"\n         0 - [a-z\\d]",
			"\n         1 - [a-zA-Z\\d]",
			"\n         2 - [\\w\\d]",
			"\n         3 - [\\w\\d!\"§$%&/()=?-]",
			"\n         4 - .",
			"\n  -rx   Regex of matched chars used to form",
			"\n        password (defaultly '[\\w]')",
		)
		os.Exit(0)
	}

	var (
		pwlen  = 32
		number = 1
		regex  = regexp.MustCompile(`[\w]`)
	)

	if pwlenres := getArgVal(args, "-l"); pwlenres != "" {
		var err error
		pwlen, err = strconv.Atoi(pwlenres)
		if err != nil {
			panic("Can not convert entered length value to int")
		}
	}

	if numberraw := getArgVal(args, "-n"); numberraw != "" {
		var err error
		number, err = strconv.Atoi(numberraw)
		if err != nil {
			panic("Can not convert entered number value to int")
		}
	}

	if strengthres := getArgVal(args, "-s"); strengthres != "" {
		strength, err := strconv.Atoi(strengthres)
		if err != nil {
			panic("Can not convert entered strength value to int")
		}
		switch strength {
		case 0:
			regex = regexp.MustCompile(STRENGTH_0)
		case 1:
			regex = regexp.MustCompile(STRENGTH_1)
		case 2:
			regex = regexp.MustCompile(STRENGTH_2)
		case 3:
			regex = regexp.MustCompile(STRENGTH_3)
		case 4:
			regex = regexp.MustCompile(STRENGTH_4)
		}
	}

	if regexres := getArgVal(args, "-rx"); regexres != "" {
		regex = regexp.MustCompile(regexres)
	}

	fullcharset := createcharset()

	for i := 0; i < number; i++ {
		fmt.Println(genPasswd(pwlen, fullcharset, regex))
	}
}
