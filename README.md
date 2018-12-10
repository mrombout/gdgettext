# gdtexttext

[![Build Status](https://travis-ci.com/mrombout/gdgettext.svg?branch=master)](https://travis-ci.com/mrombout/gdgettext) [![Go Report Card](https://goreportcard.com/badge/github.com/mrombout/gdgettext)](https://goreportcard.com/report/github.com/mrombout/gdgettext)

The `gdgettext` program extracts translatable strings from given Godot input files. It is meant as an alternative for `xgettext` specifically for games using [Godot Engine](https://godotengine.org/).

Currently it extracts translatable strings using the following rules:

* `.tscn` - every `text` parameter being set
* `.gd` - every call to `tr()`

## Installation

    $ go get github.com/mrombout/gdgettext

## Usage

    $ gdgettext [option] [inputfile]...

### Options

```
-o file
--output=file
    Write output to specified file (instead of `messages.po`).
```

### Poedit

Go to _File - Preferences_ and add a new extractor in the _Extractors_ tab and enter the following settings:

| Setting                                                     | Value              |
| ----------------------------------------------------------- | ------------------ |
| List of extensions separated by semicolons (e.g. *.cpp;*.h) | ‪*.gd;*.tscn        |
| Command to extract translations                             | ‪gdgettext -o %o %F |
| An item in input files list                                 | %f                 |
travis g
### Godot

Use `tr()` in your `.gd` files or set the `text` attributes on the `Control`s in your scene (e.g. `Button`) to make strings translatable.

Using poedit **gdgettext** parses all your `.tscn` and `.gd` files and creates a `.po` file containing all translatable strings.

Now use poedit to translate all translatable string to your desired messages.

You can also create new translation files from the `.po` files for each language you want to support. For each language you generally create a `.po` file using the language code as the file name (e.g. `nl.po`, `en.po`, `no.po`, `ja.po`).

Load all `.po` files in your Godot project settings.
