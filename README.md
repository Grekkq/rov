# ROV, RocksDB viewer in Go
To use you unfortunately have to build it yourself and is kind of pain in the ass as you need to have rocksdb itself. Please follow [https://github.com/facebook/rocksdb/blob/master/INSTALL.md](https://github.com/facebook/rocksdb/blob/master/INSTALL.md) and then build using:
```bash
CGO_CFLAGS="-I/path/to/rocksdb/include" \
CGO_LDFLAGS="-L/path/to/rocksdb -lrocksdb -lstdc++ -lm -lz -lsnappy -llz4 -lzstd" \
  go build
```
# Usage
To list all keys in database use:
`./rov get-all -path /path/to/database`

To get certain key from database use:
`./rov get -path /path/to/database -get key-to-get`