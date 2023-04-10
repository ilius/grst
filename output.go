package grst

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (self *RstBuilder) Write(fn string) error {
	lines, err := self.GetLines()
	if err != nil {
		return err
	}

	dirName := filepath.Dir(fn)
	err = os.MkdirAll(dirName, 0o755)
	if err != nil {
		return err
	}
	log.Println("created directory:", dirName)

	file, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	var numBytes int
	for _, line := range lines {
		nb, err := fmt.Fprintln(w, line)
		numBytes += nb
		log.Println(err)
	}

	err = w.Flush()
	if err != nil {
		log.Println(err)
	}

	log.Printf("wrote %d bytes to file '%s'.", numBytes, fn)

	return nil
}

func (self *RstBuilder) Print() error {
	lines, err := self.GetLines()

	fmt.Println(strings.Join(lines, "\n"))

	return err
}

func (self *RstBuilder) Resolve() (string, error) {
	lines, err := self.GetLines()

	return strings.Join(lines, "\n") + "\n", err
}
