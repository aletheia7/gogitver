// +build mage

package main

import (
	"bytes"
	"fmt"
	"github.com/magefile/mage/sh"
	"io/ioutil"
	"os"
)

// gogitver
func Gogitver() error {
	git_version, err := sh.Output("git", "describe", "--tags", "--long")
	if err != nil {
		return err
	}
	proposed_git_describe := []byte(fmt.Sprintf("package gogitver\nconst git_describe = `%v`\n", git_version))
	git_describe_fn := `vendor/gogitver/git_describe.go`
	git_describe, err := ioutil.ReadFile(git_describe_fn)
	if os.IsNotExist(err) || !bytes.Equal(proposed_git_describe, git_describe) {
		fmt.Println("new git tag:", git_version)
		if err = ioutil.WriteFile(git_describe_fn, proposed_git_describe, 0666); err != nil {
			return err
		}
	}
	return nil
}
