# GoSum

A small executable which allows you to verify or generate checksum for one, or multiple files

## How to use

- `-a {algos}` select which hash methods to use
    - `default: all`
    - available `crc32, md5, sha1, sha256, sha384, sha512`
    - Examples: `-a crc32`, `-a all`, `-a md5,sha1,sha512`
    
- `-f {files}` select which files to use
    - Examples: `-f foo.txt`, `-f foo.txt,bar.txt`, `~/*`

- `-v {hash to validate}` checks if the files have the same hash ( supports partial )
    - prefixing your partial hash by `-` will try to match the end of the hash
    - Examples: `-v fec0d`, `-v -fec0d`, `-v 0800fc577294c34e0b28ad2839435945`
 
 Example real world usage:
    
   - `gosum -a sha256 -f ubuntu-20.04.iso -v -9a9264df9f`
   - `gosum -a md5,sha256,sha512 -f my_distributable_exe`
    
## How to build

You will need [Go lang](https://golang.org/) to build the code.

Open a terminal in the directory and do `go build main.go hasher.go arguments.go verifier.go`