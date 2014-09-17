// Copyright 2014 aletheia7.
//
// This file is part of gogitver.
//
// gogitver is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either gogitver 3 of the License, or
// (at your option) any later gogitver.
//
// gogitver is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with gogitver.  If not, see <http://www.gnu.org/licenses/>.

// Package gogitver embeds a git tag into your application
//
// Available make targets:
//
//  make
//  make strip
//  make clean
//
package gogitver

import "strings"

// Tag returns a git tag
//
// Tag is the left hand side of -g
// Example: 1.0.2
func Tag() string {
	return strings.Split(git_describe, "-")[0]
}

// Git returns a "git describe --tags --long"
//
// Example: 1.0.2-0-g641e39f
func Git() string {
	return git_describe
}
