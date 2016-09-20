// Copyright 2016 aletheia7. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause
// license that can be found in the LICENSE file.
// Package gogitver embeds a git tag into your application

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
