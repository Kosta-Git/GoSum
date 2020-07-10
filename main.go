package main

import (
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"
	"sync"
)

type CheckSum struct {
	file string
	algo string
}

var lock sync.WaitGroup

func main() {
	arguments := ParseArguments()

	for _, f := range arguments.Files {
		for _, a := range arguments.Algorithms {
			lock.Add(1)
			go checkSum(
				CheckSum{f, a},
				arguments.ShouldVerify,
				arguments.Verify,
				arguments.InvertedVerification,
			)
		}
	}

	lock.Wait()
	fmt.Println("All hashes were computed")
}

func checkSum(checkSum CheckSum, shouldVerify bool, toVerify string, inverted bool) {
	defer lock.Done()

	hasher := HasherFactory(checkSum.algo)

	hash, err := hashFile(checkSum.file, hasher)
	if err != nil {
		fmt.Println("Could not compute hash")
		fmt.Println(err)
		fmt.Println("Exiting...")
		os.Exit(1)
	} else {
		fmt.Print(ensureShort(checkSum.file) + " " + checkSum.algo + ": ")

		if shouldVerify {
			verified := Verify(
				hex.EncodeToString(hash),
				toVerify,
				inverted,
			)

			PrintVerified(verified)
			fmt.Println()
		} else {
			fmt.Println(hex.EncodeToString(hash))
		}
	}
}

func hashFile(filename string, h hash.Hash) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(h, f)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func ensureShort(filename string) string {
	if strings.Contains(filename, "/") {
		fullname := strings.Split(filename, "/")

		return fullname[len(fullname)-1]
	}

	return filename
}
