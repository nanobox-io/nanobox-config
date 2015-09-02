// -*- mode: go; tab-width: 2; indent-tabs-mode: 1; st-rulers: [70] -*-
// vim: ts=4 sw=4 ft=lua noet
//--------------------------------------------------------------------
// @author Daniel Barney <daniel@nanobox.io>
// @copyright 2015, Pagoda Box Inc.
// @doc
//
// @end
// Created :   12 August 2015 by Daniel Barney <daniel@nanobox.io>
//--------------------------------------------------------------------
package config

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var (
	Config = map[string]string{}
	loaded = false
)

func Load(defaults map[string]string, file string) error {
	if loaded {
		return errors.New("config file already loaded")
	}
	loaded = true

	// load in the defaults
	for key, value := range defaults {
		Config[key] = value
	}
	if file != "" {
		// load the config file
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		scan := bufio.NewScanner(f)

		for scan.Scan() {
			line := scan.Text()
			switch {
			case strings.HasPrefix(line, "#"): // comment
			case len(line) == 0: // empty line
			default:
				key, value, err := parseArg(line)
				if err != nil {
					return err
				}
				Config[key] = value
			}
		}
	}

	// load in the cli arguments
	args := os.Args[1:]
	for _, arg := range args {
		key, value, err := check(arg)
		if err == nil {
			Config[key] = value
		}
	}
	return nil
}

func check(arg string) (string, string, error) {
	if strings.HasPrefix(arg, "-") {
		arg = strings.TrimPrefix(arg, "-")
		return parseArg(arg)
	}
	return "", "", errors.New("not a config parameter")
}

func parseArg(arg string) (string, string, error) {
	fields := strings.SplitN(arg, " ", 2)
	if len(fields) != 2 {
		return "", "", errors.New("bad config file")
	}
	return fields[0], fields[2], nil
}
