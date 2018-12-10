package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMainCall(t *testing.T) {
	outputFileName := "test/test_output.po"
	expectedMsgID := `msgid "Hello World!"`

	err := os.RemoveAll(outputFileName)
	if err != nil {
		t.Fatal(err)
	}

	os.Args = []string{
		"gdgettext",
		"-o",
		outputFileName,
		"test/test_input.gd",
	}
	main()

	outputFileContents, err := ioutil.ReadFile(outputFileName)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(outputFileContents), expectedMsgID) {
		t.Fatalf("expected %v to contain %v", outputFileName, expectedMsgID)
	}
}
