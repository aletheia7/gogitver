// Copyright 2016 aletheia7. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause
// license that can be found in the LICENSE file.

// Package gogitver embeds a git tag into your application
//
// Available make targets:
//
//  make
//  make strip
//  make clean
//
package gogitver

import (
	"regexp"
)

// Tag returns a git tag
//
// Tag is the left hand side of 1.0.2-α-0-g1234567
// Example: 1.0.2-α
func Tag() string {
	re := regexp.MustCompile(`^(.+)-\d+-g[[:xdigit:]]{7}$`)
	return re.FindStringSubmatch(git_describe)[1]
}

// Git returns a "git describe --tags --long"
//
// Example: 1.0.2-0-g641e39f
func Git() string {
	return git_describe
}
