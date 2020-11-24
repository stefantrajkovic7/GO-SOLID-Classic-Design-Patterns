package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0
type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// ...
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}


func main() {
    j := Journal{}
    j.AddEntry("Hello World")
    j.AddEntry("I ate a bug")
    fmt.Println(j.String())

    j.Save("Journal.txt")
}