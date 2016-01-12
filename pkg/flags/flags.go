// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2015 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

// Package flags handles the given flags
package flags

import (
	"flag"
	"strconv"
)

// Options has fields for the existing flags
type Options struct {
	Action, Release, Channel, Arch, LogLevel, Qcow2compat string
}

const (
	defaultAction      = "create"
	defaultRelease     = "rolling"
	defaultChannel     = "edge"
	defaultArch        = "amd64"
	defaultLogLevel    = "info"
	defaultQcow2compat = "1.1"
)

// Parse analyzes the flags and returns a Options instance with the values
func Parse() *Options {
	var (
		action      = flag.String("action", defaultAction, "action to be performed")
		release     = flag.String("release", defaultRelease, "release of the image to be created")
		channel     = flag.String("channel", defaultChannel, "channel of the image to be created")
		arch        = flag.String("arch", defaultArch, "arch of the image to be created")
		logLevel    = flag.String("loglevel", defaultLogLevel, "Level of the log putput, one of debug, info, warning, error, fatal, panic")
		qcow2compat = flag.String("qcow2compat", defaultQcow2compat, "Qcow2 compatibility level (0.10 or 1.1)")
	)
	flag.Parse()
	dotRelease := addDot(*release)
	return &Options{
		Action: *action, Release: dotRelease, Channel: *channel, Arch: *arch, LogLevel: *logLevel, Qcow2compat: *qcow2compat}
}

func addDot(release string) string {
	if len(release) == 4 {
		if _, err := strconv.Atoi(release); err == nil {
			return release[0:2] + "." + release[2:]
		}
	}
	return release
}
