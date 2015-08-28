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
	"testing"
)

func TestParser(test *testing.T) {
	// this test really needs to be finished
	defaults := map[string]string{
		"key": "value"}
	Load(defaults, "")

}
