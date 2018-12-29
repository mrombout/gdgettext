package main

import (
	"io"
	"text/template"
)

const poTemplate = `# SOME DESCRIPTIVE TITLE.
# Copyright (C) YEAR THE PACKAGE'S COPYRIGHT HOLDER
# This file is distributed under the same license as the PACKAGE package.
# FIRST AUTHOR <EMAIL@ADDRESS>, YEAR.
#
#, fuzzy
msgid ""
msgstr ""
"Project-Id-Version: PACKAGE VERSION\n"
"Report-Msgid-Bugs-To: \n"
"POT-Creation-Date: 2018-12-09 18:30+0100\n"
"PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE\n"
"Last-Translator: FULL NAME <EMAIL@ADDRESS>\n"
"Language-Team: LANGUAGE <LL@li.org>\n"
"Language: \n"
"MIME-Version: 1.0\n"
"Content-Type: text/plain; charset=UTF-8\n"
"Content-Transfer-Encoding: 8bit\n"
{{range .Translations}}
#: {{range $index, $loc := .Files}}{{if $index}} {{end}}{{$loc.File}}:{{$loc.Line}}{{end}}
msgid "{{.Msgid}}"
msgstr "{{.Msgstr}}"
{{end}}`

func render(poFile poFile, writer io.Writer) error {
	tmpl, err := template.New("default").Parse(poTemplate)
	if err != nil {
		return err
	}

	err = tmpl.Execute(writer, poFile)
	if err != nil {
		return err
	}

	return nil
}
