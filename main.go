package main

import (
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"
)

func main() {
	arguments := ParseArguments()

	for _, f := range arguments.Files {
		for _, a := range arguments.Algorithms {
			hasher := HasherFactory(a)

			hash, err := hashFile(f, hasher)
			if err != nil {
				fmt.Println("Could not compute hash")
				fmt.Println(err)
				fmt.Println("Exiting...")
				os.Exit(1)
			} else {
				fmt.Print(ensureShort(f) + " " + a + ": ")

				if arguments.ShouldVerify {
					verified := Verify(
						hex.EncodeToString(hash),
						arguments.Verify,
						arguments.InvertedVerification,
					)

					PrintVerified(verified)
					fmt.Println()
				} else {
					fmt.Println(hex.EncodeToString(hash))
				}
			}
		}

		fmt.Println()
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
