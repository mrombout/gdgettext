# gdtexttext

> [!IMPORTANT]
> Since Godot 4.0, the editor can generate a PO template automatically from specified scene and GDscript files.
> See [Automatic generation using the editor](https://docs.godotengine.org/en/4.2/tutorials/i18n/localization_using_gettext.html#automatic-generation-using-the-editor) for more information.
> You don't need `gdgettext` anymore.

[![Actions Status](https://github.com/mrombout/gdgettext/workflows/Main/badge.svg)](https://github.com/mrombout/gdgettext/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/mrombout/gdgettext)](https://goreportcard.com/report/github.com/mrombout/gdgettext)

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

Open _Poedit Preferences_ and add a new extractor in the _Extractors_ tab and enter the following settings:

| Setting                                                     | Value              |
| ----------------------------------------------------------- | ------------------ |
| Language                                                    | Godot              |
| List of extensions separated by semicolons (e.g. *.cpp;*.h) | ‪*.gd;*.tscn        |
| Command to extract translations                             | ‪gdgettext -o %o %F |
| An item in input files list                                 | %f                 |

Then create a new `.po` file and press _Extract from sources_ and set the root of your Godot project as the `Base path` and add paths containing the translatable strings to the `Paths` section.

After this Poedit will use `gdgettext` to extract all translatable strings and you can start editing the `.po` file using Poedit.

### Godot

Use `tr()` in your `.gd` files or set the `text` attributes on the `Control`s in your scene (e.g. `Button`) to make strings translatable.

Using poedit **gdgettext** parses all your `.tscn` and `.gd` files and creates a `.po` file containing all translatable strings.

Now use poedit to translate all translatable string to your desired messages.

You can also create new translation files from the `.po` files for each language you want to support. For each language you generally create a `.po` file using the language code as the file name (e.g. `nl.po`, `en.po`, `no.po`, `ja.po`).

Load all `.po` files in your Godot project settings.
