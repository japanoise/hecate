package main

import (
	"encoding/json"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/nsf/termbox-go"
)

type Style struct {
	Default_bg termbox.Attribute
	Default_fg termbox.Attribute

	Rune_fg       termbox.Attribute
	Space_rune_fg termbox.Attribute
	Int_fg        termbox.Attribute
	Bit_fg        termbox.Attribute

	Selected_option_bg termbox.Attribute
	Search_progress_fg termbox.Attribute

	Text_cursor_hex_bg termbox.Attribute
	Bit_cursor_hex_bg  termbox.Attribute
	Int_cursor_hex_bg  termbox.Attribute
	Fp_cursor_hex_bg   termbox.Attribute

	Hilite_hex_fg  termbox.Attribute
	Hilite_rune_fg termbox.Attribute

	Field_editor_bg termbox.Attribute
	Field_editor_fg termbox.Attribute

	Field_editor_last_bg termbox.Attribute
	Field_editor_last_fg termbox.Attribute

	Field_editor_invalid_bg termbox.Attribute
	Field_editor_invalid_fg termbox.Attribute

	About_logo_bg termbox.Attribute

	Filled_bit_rune rune
	Empty_bit_rune  rune
	Space_rune      rune
}

func (s *Style) SaveStyleToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(s)
}

func LoadStyleFromFile(filename string) (*Style, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	ret := &Style{}
	err = decoder.Decode(ret)
	return ret, err
}

func GetConfigDir() (string, error) {
	configdir := os.Getenv("XDG_CONFIG_HOME")
	if configdir == "" {
		h, err := homedir.Dir()
		if err != nil {
			return configdir, err
		}
		configdir = h + string(os.PathSeparator) + ".config" + string(os.PathSeparator) + PROGRAM_NAME
	} else {
		configdir += string(os.PathSeparator) + PROGRAM_NAME
	}
	return configdir + string(os.PathSeparator), os.MkdirAll(configdir, 0755)
}
