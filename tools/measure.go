package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readLines(path string) []string {
	srcBytes, err := ioutil.ReadFile(path)
	check(err)
	return strings.Split(string(srcBytes), "\n")
}

var commentPat = regexp.MustCompile("\\s*\\/\\/")

func main() {
	sourcePaths, err := filepath.Glob("./examples/*/*")
	check(err)
	foundLongFile := false
	for _, sourcePath := range sourcePaths {
		foundLongLine := false
		lines := readLines(sourcePath)
		for i, line := range lines {
			// Convert tabs to spaces before measuring, so we get an accurate measure
			// of how long the output will end up being.
			line := strings.Replace(line, "\t", "    ", -1)
			count := utf8.RuneCountInString(line)
			if !foundLongLine && !commentPat.MatchString(line) && (count > 58) {
				fmt.Printf("measure: %s:%d %d\n", sourcePath, i+1, count)
				foundLongLine = true
				foundLongFile = true
			}
		}
	}
	if foundLongFile {
		os.Exit(1)
	}
}
