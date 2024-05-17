package main

import (
	"bufio"
	"os"
	"path/filepath"
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
		f, e := os.Open(filePath)
		if e != nil {
			panic(e)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		line := 0
		s, e := Readln(r)
		for e == nil {
			line++
			linestr := s
			keys := extract(linestr)
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
			s, e = Readln(r)
		}

	}

	return poFile, nil
}
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
func isSupportedExt(ext string) bool {
	for _, supportedExt := range supportedExts {
		if ext == supportedExt {
			return true
		}
	}

	return false
}
