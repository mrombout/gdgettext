package main

import (
	"flag"
	"os"
)

var outputFile string

const (
	outputFlagDefault = "messages.po"
	outputFlagUsage   = "Write output to specified file (instead of `messages.po`)."
)

func init() {
	flag.StringVar(&outputFile, "o", outputFlagDefault, outputFlagUsage)
	flag.StringVar(&outputFile, "output", outputFlagDefault, outputFlagUsage)
}

func main() {
	flag.Parse()
	files := flag.Args()

	poFile, err := process(osFileSystem{}, files)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	err = render(poFile, outputFile)
	if err != nil {
		panic(err)
	}
}
