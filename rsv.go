package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/linxGnu/grocksdb"
)

func handleGet(path string, key string) {
	ro := grocksdb.NewDefaultReadOptions()
	defer ro.Destroy()
	opts := grocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(false)
	defer opts.Destroy()

	db, err := grocksdb.OpenDbForReadOnly(opts, path, true)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer db.Close()

	value, err := db.Get(ro, []byte(key))
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer value.Free()
	if !value.Exists() {
		fmt.Printf("There is no entry for key: %v (raw bytes: %v) in database: %s\n", key, []byte(key), path)
		return
	}
	fmt.Printf("Raw data: %v\n", value.Data())
	fmt.Printf("Data parsed to string: %s\n", value.Data())
}

func handleGetAll(path string) {
	ro := grocksdb.NewDefaultReadOptions()
	defer ro.Destroy()
	opts := grocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(false)
	defer opts.Destroy()

	db, err := grocksdb.OpenDbForReadOnly(opts, path, true)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer db.Close()

	it := db.NewIterator(ro)
	defer it.Close()
	it.SeekToFirst()
	for it = it; it.Valid(); it.Next() {
		key := it.Key()
		value := it.Value()
		fmt.Printf("Parsed key: %s raw bytes: %v\n", key.Data(), key.Data())
		key.Free()
		value.Free()
	}
	if err := it.Err(); err != nil {
		fmt.Errorf(err.Error())
	}
}

func main() {
	getCommand := flag.NewFlagSet("get", flag.ExitOnError)
	getAllCommand := flag.NewFlagSet("get-all", flag.ExitOnError)

	rawPath := getCommand.String("path", ".", "Path that contains database")
	rawKey := getCommand.String("get", "", "Key to extract from database (provided string will be casted to bytes)")
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
		key := ""
		if rawKey != nil {
			key = *rawKey
		}
		fmt.Printf("Using database in: %s, getting key: %s\n", *rawPath, key)
		handleGet(path, key)
	case "get-all":
		getAllCommand.Parse(os.Args[2:])
		if rawPathAll != nil {
			path = *rawPathAll
		}
		fmt.Printf("Using database in: %s, getting all-keys\n", path)
		handleGetAll(path)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
