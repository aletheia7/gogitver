git_cmd = git describe --tags --long
# Uncomment for tag without git hash
#git_cmd = git describe --tags

verbose_go_cmds = -v
strip_options = --strip-all
package = $(notdir $(PWD))
git_describe_go = gogitver/git_describe.go
go_code = printf "package gogitver\nconst git_describe = \`%s\`\\n" `$(git_cmd)`

all : go_install

go_install : remove_older_describe_go $(git_describe_go) 
	go install $(verbose_go_cmds) $(package)

# go install followed by strip to remove symbols
strip : all
	$(shell strip $(strip_options) `go env GOPATH`/bin/$(package))

# make git_describe_go with git tag
$(git_describe_go) :
	$(shell $(go_code) > $@)

# Compare git tag with current tag
# If git tag is newer than existing tag, remove git_describe_go and remake 
remove_older_describe_go : $(git_describe_go) force 
	@printf '%s' $(shell $(go_code) | diff -q - $(git_describe_go) >/dev/null ; if [ $$? -ne 0 ] ; then (rm $(git_describe_go) ; printf "new git tag: %s\n" `$(git_cmd)`;  $(MAKE) $(MAKEFLAGS) $(git_describe_go) > /dev/null ) ; fi)

# git clean
.PHONY : clean
clean :
	go clean -i $(verbose_go_cmds) $(package)

.PHONY : force
force :
