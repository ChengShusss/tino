package main

import (
	"fmt"
	"os"
	"time"

	"atomicgo.dev/cursor"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/term"
)

type Info struct {
	Header string
	Body   string
	Footer string
}

func (i *Info) Sprint(lw int) string {
	var s string
	s += ExtractLine(i.Header, TermWidth, 0, 1)
	s += ExtractLine(i.Body, TermWidth, 0, TermHeight-2)
	s += ExtractLine(i.Footer, lw, 0, 1)
	return s
}

var (
	TermHeight, TermWidth int
)

func main() {

	// fmt.Print("\033[2J\033[0;0H")
	TermWidth, TermHeight, err := term.GetSize(0)
	if err != nil {
		fmt.Printf("Failed to get terminal size, err: %v\n", err)
		os.Exit(1)
	}
	// fmt.Printf("Hight: %d, witdh: %d\n", TermHeight, TermWidth)

	info := Info{}
	area := cursor.NewArea()
	for {
		v, _ := mem.VirtualMemory()

		info.Header = fmt.Sprintf("H: %d --- W: %d\n", TermHeight, TermWidth)
		info.Body = "1\n2\n3\n4\n5\n6\n" + fmt.Sprintf("Total: %vMB, Free:%vMB, UsedPercent:%.2f%%\n", v.Total/1024/1024, v.Free/1024/1024, v.UsedPercent)
		info.Footer = "\033[1;31m" + time.Now().Format("2006-01-02 15:04:05") + "\033[0m"

		// fmt.Print("\033[2J\033[0;0H")
		area.Update(info.Sprint(TermWidth))

		// fmt.Print(info.Sprint(190))
		time.Sleep(1 * time.Second)
	}

}

// Extract lines from
func ExtractLine(str string, lineWidth, start, lens int) string {
	var cnt, idxStart, idxEnd int
	// fmt.Printf("lw: %v, start: %v, lens: %v\n", lineWidth, start, lens)

	pos := 0
	for i := 0; i < len(str); i++ {
		// fmt.Printf("idx: %v, char: %d, cnt: %v\n", i, str[i], cnt)
		if str[i] == '\033' {
			for ; i < len(str) && str[i] != 'm'; i++ {
			}
			continue
		}

		pos += 1
		if str[i] == '\n' || pos >= lineWidth {
			cnt += 1
			pos = 0
			if cnt == start {
				// fmt.Printf("Set idxStart at %d\n", i)
				idxStart = i + 1
			}
			if cnt == start+lens {
				// fmt.Printf("Set idxEnd at %d\n", i)
				idxEnd = i
				break
			}
		}
	}
	if idxEnd <= idxStart {
		// fmt.Printf("Go into here to set idxEnd\n")
		idxEnd = len(str)
	}

	// fmt.Printf("Start: %v, end: %v, s:[%s]\n", idxStart, idxEnd, str)
	// fmt.Printf("%v\n", []byte(str))

	return str[idxStart:idxEnd]
}
