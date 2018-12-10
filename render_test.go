package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestRender(t *testing.T) {
	expectedOutput, err := ioutil.ReadFile("test/test_render.po")
	if err != nil {
		t.Fatal(err)
	}

	poFile := poFile{
		Translations: map[msgID]message{
			"test": {
				Files: []messageLocation{
					{
						File: "my_script.gd",
						Line: 1,
					},
					{
						File: "my_script.tscn",
						Line: 76,
					},
				},
				Msgid:  "test",
				Msgstr: "Hello World!",
			},
		},
	}

	buf := bytes.Buffer{}

	err = render(poFile, &buf)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(buf.Bytes(), expectedOutput) {
		t.Fatalf("output does not match expected output\n\n ACTUAL OUTPUT ====\n%v\nEXPECTED OUTPUT ====\n%v", buf.String(), string(expectedOutput))
	}
}
