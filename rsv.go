package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/linxGnu/grocksdb"
)

func main() {
	getCommand := flag.NewFlagSet("get", flag.ExitOnError)
	getAllCommand := flag.NewFlagSet("get-all", flag.ExitOnError)

	rawPath := getCommand.String("path", ".", "Path that contains database")
	rawKey := getCommand.String("get", "", "Key to extract from database")
	rawPathAll := getAllCommand.String("path", ".", "Path that contains database")
	path := "."

	if len(os.Args) < 2 {
		fmt.Println("specify either get or get-all")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "get":
		getCommand.Parse(os.Args[2:])
		if rawPath != nil {
			path = *rawPath
		}
	case "get-all":
		getAllCommand.Parse(os.Args[2:])
		if rawPathAll != nil {
			path = *rawPathAll
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	key := ""
	if rawKey != nil {
		key = *rawKey
	}

	ro := grocksdb.NewDefaultReadOptions()
	defer ro.Destroy()

	if getCommand.Parsed() {
		fmt.Printf("Using database in: %s, getting key: %s", *rawPath, key)
		opts := grocksdb.NewDefaultOptions()
		opts.SetCreateIfMissing(false)
		defer opts.Destroy()
		db, err := grocksdb.OpenDbForReadOnly(opts, path, true)
		defer db.Close()
		// db, err := grocksdb.OpenDb(opts, path)
		if err != nil {
			fmt.Errorf(err.Error())
		}

		value, err := db.Get(ro, []byte(key))
		if err != nil {
			fmt.Errorf(err.Error())
		}
		defer value.Free()

		fmt.Printf("Exists: %v, Data: %v", value.Exists(), value.Data())
		fmt.Printf("After casting: %s", value.Data())
	}

	if getAllCommand.Parsed() {
		fmt.Printf("Using database in: %s, getting all-keys", path)
		opts := grocksdb.NewDefaultOptions()
		opts.SetCreateIfMissing(false)
		defer opts.Destroy()
		db, err := grocksdb.OpenDbForReadOnly(opts, path, true)
		defer db.Close()
		if err != nil {
			fmt.Errorf(err.Error())
		}
		it := db.NewIterator(ro)
		defer it.Close()

		it.Seek([]byte("foo"))
		for it = it; it.Valid(); it.Next() {
			key := it.Key()
			value := it.Value()
			fmt.Printf("Key: %v parsed: %s\n", key.Data(), key.Data())
			key.Free()
			value.Free()
		}
		if err := it.Err(); err != nil {
			fmt.Errorf(err.Error())
		}

	}
}
