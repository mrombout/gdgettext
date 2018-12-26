package main

import "testing"

func TestExtractFromGDScript(t *testing.T) {
	expectedTokens := []string{
		"Hello,",
		"World!",
	}
	expectedNumTokensFound := len(expectedTokens)
	content := `
	extends Node

	func _ready():
		some_function(tr("Hello,"))
		some_function(tr('World!'))
	`

	results := extract(content)

	actualNumTokensFound := len(results)
	if actualNumTokensFound != expectedNumTokensFound {
		t.Fatalf("expected to have extracted %v token, but found %v", expectedNumTokensFound, results)
	}

	for index := range expectedTokens {
		actualToken := results[index]
		expectedToken := expectedTokens[index]
		if actualToken != expectedToken {
			t.Fatalf("expected token '%v' (%v) to equal token '%v' (%v) in the results (%v)", index, actualToken, index, expectedToken, results)
		}
	}
}

func TestExtractEscapedStringFromGDScript(t *testing.T) {
	expectedTokens := []string{
		"It\\'s a string!",
	}
	expectedNumTokensFound := len(expectedTokens)
	content := `
	extends Node

	func _ready():
		some_function(tr('It\'s a string!'))
	`

	results := extract(content)

	actualNumTokensFound := len(results)
	if actualNumTokensFound != expectedNumTokensFound {
		t.Fatalf("expected to have extracted %v token, but found %v", expectedNumTokensFound, results)
	}

	for index := range expectedTokens {
		actualToken := results[index]
		expectedToken := expectedTokens[index]
		if actualToken != expectedToken {
			t.Fatalf("expected token '%v' (%v) to equal token '%v' (%v) in the results (%v)", index, actualToken, index, expectedToken, results)
		}
	}
}
