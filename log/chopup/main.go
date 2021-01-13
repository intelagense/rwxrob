package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// just for this specific parsing need, doesn't convey
func parseTime(s string) (time.Time, error) {
	f := "Monday, January 2, 2006, 3:04:05PM"
	return time.ParseInLocation(f, s, time.Local)
}

type Entry struct {
	datetime time.Time
	markdown string
}

func (e Entry) String() string {
	f := "Monday, January 2, 2006, 3:04:05PM MST"
	return fmt.Sprintf("## %v [%v]\n%v",
		e.datetime.Format(f),
		e.datetime.Unix(),
		e.markdown,
	)
}

var curEntry = &Entry{}

var entries = []*Entry{}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var inCodeBlock bool
	for s.Scan() {
		line := s.Text()

		if strings.HasPrefix(line, "```") {
			curEntry.markdown += line + "\n"
			if inCodeBlock {
				inCodeBlock = false
				continue
			}
			inCodeBlock = true
			continue
		}

		if inCodeBlock {
			curEntry.markdown += line + "\n"
			continue
		}

		if strings.HasPrefix(line, "## ") {
			line = strings.TrimLeft(line, "# ")
			t, err := parseTime(line)
			if err != nil {
				fmt.Println(err)
				return
			}
			entries = append(entries, curEntry)
			curEntry = new(Entry)
			curEntry.datetime = t
			continue
		}

		curEntry.markdown += line + "\n"

	}

	for _, entry := range entries[1:] {
		p := filepath.Join(
			fmt.Sprintf("%v", int(entry.datetime.Year())),
			fmt.Sprintf("%02v", int(entry.datetime.Month())),
			fmt.Sprintf("%02v", int(entry.datetime.Day())),
		)
		readme := filepath.Join(p, "README.md")
		fmt.Println(readme)
		fmt.Print(entry)
	}

	// TODO scan lines in standard input
	// for each line:
	//   if in code block just continue
	//   if line starts with code block then
	//     set inblock
	//     continue
	//   if line starts with one or more # and next line is blank then
	//     parse heading line
	//     close and save previous block
	//     open a new block
	//     set the heading text for the new block
	//     continue
}
