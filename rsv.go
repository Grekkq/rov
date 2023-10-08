package main

import (
	"flag"
	"fmt"

	"github.com/linxGnu/grocksdb"
)

func main() {
	rawPath := flag.String("path", ".", "Path that contains database")
	flag.Parse()
	path := "."
	if rawPath != nil {
		path = *rawPath
	}
	fmt.Printf("Using database in: %s", path)
	grocksdb.OpenDbForReadOnly(nil, "", true)
	// db, err := rocksdb.OpenDBReadOnly(path)
	// if err != nil {
	// 	fmt.Errorf(err.Error())
	// }
	// db.GetProperty("")
	// asd, err := rocksdb.CreateDB("")
}
