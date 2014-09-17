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

package gogitver_test

import (
	"fmt"
	"gogitver"
	"testing"
)

func ExampleGit() {
	fmt.Println("git:", gogitver.Git())
	// git: 1.0.2-0-g641e39f
}

func ExampleTag() {
	fmt.Println("tag:", gogitver.Tag())
	// tag: 1.0.2
}

func TestGit(t *testing.T) {
	t.Log(gogitver.Git())
}

func TestTag(t *testing.T) {
	t.Log(gogitver.Tag())
}
