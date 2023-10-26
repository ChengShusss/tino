package main

import (
	"fmt"
	"testing"
	"time"
)

func TestExtractLine(t *testing.T) {
	s := "1234567890\n\naaaa\n\n"
	s = "1\n2\n3\n4\n5\n6\n"

	fmt.Println(ExtractLine(s, 190, 1, 3))
}

func TestGenContent(t *testing.T) {
	info := Info{}

	TermHeight = 10
	TermWidth = 190

	info.Header = fmt.Sprintf("H: %d --- W: %d\n", TermHeight, TermWidth)
	info.Body = "1\n2\n3\n4\n5\n6\n"
	info.Footer = time.Now().Format("20060102150405")

	fmt.Println(info.Sprint(10))

	// fmt.Printf("%s===\n", info.Body)
	// fmt.Println(ExtractLine(info.Body, TermWidth, 0, TermHeight-2))
	fmt.Println((time.Now().Format("\033[1;31;40m2006-01-02\033[0m 15:04:05")))
}

func TestColorInfo(t *testing.T) {
	s := (time.Now().Format("\033[1;31;40m2006-01-02\033[0m 15:04:05"))

	res := ExtractLine(s, 190, 0, 1)

	fmt.Println(res)
}
