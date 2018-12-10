package main

import (
	"errors"
	"testing"
)

var errUnreadable = errors.New("can not read file")

type mockFileSystem struct {
}

func (mockFileSystem) readFile(filename string) ([]byte, error) {
	if filename == "unsupported_file.csv" {
		return []byte(`this;is;a;nice;csvtr("tr")`), nil
	} else if filename == "unreadable_file.tscn" {
		return []byte{}, errUnreadable
	} else if filename == "single_translatable.gd" {
		return []byte(`extends Node

func _ready():
	my_function(tr("Hello!"))
`), nil
	} else if filename == "single_translatable.tscn" {
		return []byte(`[gd_scene load_steps=7 format=2]

[node name="start_button" type="Button" parent="." index="3"]

shortcut = null
group = null
text = "Start game"
flat = false
align = 1`), nil
	} else if filename == "empty_msgid.gd" {
		return []byte(`extends Node

func _ready():
	my_function(tr(""))
`), nil
	} else if filename == "duplicate_msgids.gd" {
		return []byte(`extends Node

func _ready():
	my_function(tr("new game"), tr("new game"))
`), nil
	}

	return []byte{}, nil
}

func TestProcessNoFilePaths(t *testing.T) {
	_, err := process(mockFileSystem{}, []string{})

	if err != nil {
		t.Fatalf("expected error to be nil, got '%v'", err)
	}
}

func TestProcessUnsupportedFile(t *testing.T) {
	poFile, err := process(mockFileSystem{}, []string{
		"unsupported_file.jpg",
	})

	if err != nil {
		t.Fatalf("expected error to be nil, got '%v'", err)
	}

	numTranslations := len(poFile.Translations)
	if numTranslations > 0 {
		t.Fatalf("expected process to skip all unsupported files resulting in no translations, but found %v (%v)", numTranslations, poFile)
	}
}

func TestProcessUnreadableFile(t *testing.T) {
	_, err := process(mockFileSystem{}, []string{
		"unreadable_file.tscn",
	})

	if err == nil {
		t.Fatalf("expected an error to be returned, but got nil")
	}
	if err != errUnreadable {
		t.Fatalf("expected returned error to be '%v', but got '%v'", errUnreadable, err)
	}
}

func TestProcessTranslatableInGDScript(t *testing.T) {
	poFile, err := process(mockFileSystem{}, []string{
		"single_translatable.gd",
	})

	if err != nil {
		t.Fatalf("expected error to be nil, got '%v'", err)
	}
	numTranslations := len(poFile.Translations)
	if numTranslations != 1 {
		t.Fatalf("expected process to find 1 translatable, but found %v (%v)", numTranslations, poFile)
	}
	if _, ok := poFile.Translations["Hello!"]; !ok {
		t.Fatalf("expected process to find msgid `Hello!`, but found only %v", poFile.Translations)
	}
}

func TestProcessTranslatableInTscn(t *testing.T) {
	poFile, err := process(mockFileSystem{}, []string{
		"single_translatable.tscn",
	})

	if err != nil {
		t.Fatalf("expected error to be nil, got '%v'", err)
	}
	numTranslations := len(poFile.Translations)
	if numTranslations != 1 {
		t.Fatalf("expected process to find 1 translatable, but found %v (%v)", numTranslations, poFile)
	}
	if _, ok := poFile.Translations["Start game"]; !ok {
		t.Fatalf("expected process to find msgid `Start game`, but found only %v", poFile.Translations)
	}
}

func TestProcessSkipEmptyMsgId(t *testing.T) {
	poFile, err := process(mockFileSystem{}, []string{
		"empty_msgid.gd",
	})

	if err != nil {
		t.Fatalf("expected error to be nil, got '%v'", err)
	}
	numTranslations := len(poFile.Translations)
	if numTranslations > 0 {
		t.Fatalf("expected process to skip all empty msgids resulting in no translations, but found %v (%v)", numTranslations, poFile)
	}
}

func TestProcessDuplicateMsgIds(t *testing.T) {
	poFile, err := process(mockFileSystem{}, []string{
		"duplicate_msgids.gd",
	})

	if err != nil {
		t.Fatalf("expected error to be nil, got '%v'", err)
	}
	numTranslations := len(poFile.Translations)
	if numTranslations != 1 {
		t.Fatalf("expected process to find 1 translatable, but found %v (%v)", numTranslations, poFile)
	}
	if _, ok := poFile.Translations["new game"]; !ok {
		t.Fatalf("expected process to find msgid `new game`, but found only %v", poFile.Translations)
	}
}
