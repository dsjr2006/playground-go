package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/golang/snappy"
	"github.com/mholt/archiver"
)

func main() {
	origin := "/Users/dsjr2006/Dev/golang/src/github.com/dsjr2006/playground-go/bolty/my.db"
	target := "/Users/dsjr2006/Downloads"
	reader, err := os.Open(origin)
	if err != nil {
		os.Exit(1)
	}
	filename := filepath.Base(origin)
	target = filepath.Join(target, fmt.Sprintf("%s.snpy", filename))
	writer, err := os.Create(target)
	if err != nil {
		os.Exit(1)
	}
	defer writer.Close()
	snap := snappy.NewBufferedWriter(writer)
	defer snap.Close()

	_, err = io.Copy(snap, reader)

	err = archiver.Zip.Make("my.db.zip", []string{origin})
}
