package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Arguments struct {
	Algorithms           []string
	Files                []string
	Verify               string
	InvertedVerification bool
	ShouldVerify         bool
}

func ParseArguments() Arguments {
	algoArg := algoArgument()
	fileArg := fileArgument()
	verifArg := verificationArgument()

	flag.Parse()

	shouldVerify := false
	isReverseVerif := false
	verifHash := ""

	if len(*verifArg) > 0 {
		if (*verifArg)[0] == '-' {
			isReverseVerif = true
			verifHash = (*verifArg)[1:]
		} else {
			isReverseVerif = false
			verifHash = *verifArg
		}

		shouldVerify = true
	}

	parsedFiles := strings.Split(*fileArg, ",")
	for i, file := range parsedFiles {
		if strings.HasPrefix(file, "~") {
			file = strings.Replace(file, "~", os.Getenv("HOME"), 1)
		}

		if strings.HasSuffix(file, "*") {
			parsedFiles = append(parsedFiles[:i], parsedFiles[i+1:]...)

			files, err := ioutil.ReadDir(strings.TrimSuffix(file, "*"))
			if err != nil {
				fmt.Println("Error parsing: ", file)
				fmt.Println(err)
				os.Exit(1)
			}

			var buffer []string
			for _, f := range files {
				if !f.IsDir() {
					fullPath := strings.TrimSuffix(file, "*") + f.Name()

					buffer = append(buffer, fullPath)
				}
			}

			temp := append(parsedFiles[:i], buffer...)
			parsedFiles = append(temp, parsedFiles[i:]...)
		}
	}

	var parsedAlgos []string
	if *algoArg == "all" {
		parsedAlgos = AvailableAlgorithms
	} else {
		parsedAlgos = strings.Split(*algoArg, ",")
	}

	return Arguments{
		Algorithms:           parsedAlgos,
		Files:                parsedFiles,
		Verify:               verifHash,
		InvertedVerification: isReverseVerif,
		ShouldVerify:         shouldVerify,
	}
}

func algoArgument() *string {
	usage := `Which algorithms to use
"all" to run all AvailableAlgorithms
available: ( ` + strings.Join(AvailableAlgorithms, ", ") + ` ) 	

Examples: "-a all", "-a md5", "-a md5,sha512"
`

	return flag.String(
		"a",
		"all",
		usage,
	)
}

func fileArgument() *string {
	usage := `Which files to check	

Examples: "-f test.txt", "-f foo.txt,bar.txt", "-f ./*,./**,~/.ssh/id_rsa"
`

	return flag.String(
		"f",
		"",
		usage,
	)
}

func verificationArgument() *string {
	usage := `Hash to verify, you can do partial verification
Prefix with "-" to check the end of the hash

Examples: "-v 995756101545048499", "-v -995756101545048499"
`

	return flag.String("v", "", usage)
}
