package main

import (
	"bufio"
	"path/filepath"
	"strings"
)

var supportedExts = []string{
	".gd",
	".tscn",
}

type msgID = string

type poFile struct {
	Translations map[msgID]message
}

type messageLocation struct {
	File string
	Line int
}

type message struct {
	Files  []messageLocation
	Msgid  msgID
	Msgstr string
}

func process(fs fileSystem, filePaths []string) (poFile, error) {
	poFile := poFile{
		Translations: map[msgID]message{},
	}

	for _, filePath := range filePaths {
		if !isSupportedExt(filepath.Ext(filePath)) {
			continue
		}

		content, err := fs.readFile(filePath)
		if err != nil {
			return poFile, err
		}

		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		line := 0
		for scanner.Scan() {
			line++

			keys := extract(string(scanner.Text()))
			for _, key := range keys {
				messageLoc := messageLocation{
					File: filepath.Base(filePath),
					Line: line,
				}
				message := message{
					Files: []messageLocation{
						messageLoc,
					},
					Msgid:  key,
					Msgstr: "",
				}

				if key == "" {
					continue
				}

				if val, ok := poFile.Translations[key]; ok {
					val.Files = append(val.Files, messageLoc)
					poFile.Translations[key] = val
				} else {
					poFile.Translations[key] = message
				}

			}
		}
	}

	return poFile, nil
}

func isSupportedExt(ext string) bool {
	for _, supportedExt := range supportedExts {
		if ext == supportedExt {
			return true
		}
	}

	return false
}
