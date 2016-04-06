git_cmd = git describe --tags --long
# Uncomment for tag without git hash
#git_cmd = git describe --tags

verbose_go_cmds = -v
strip_options = --strip-all
package = $(notdir $(PWD))
git_describe_go = $(dir $(realpath $(lastword $(MAKEFILE_LIST))))git_describe.go
go_code = printf "package gogitver\nconst git_describe = \`%s\`\\n" `$(git_cmd)`
ifeq ($(MAKECMDGOALS), strip)
go_no_debug = -ldflags '-w'
endif

all : go_install

go_install : remove_older_describe_go
	go install $(go_no_debug) $(verbose_go_cmds) $(package)

# go install followed by strip to remove symbols
strip : clean go_install
	strip $(strip_options) ../../bin/$(package)

# make git_describe_go with git tag
$(git_describe_go) :
	$(shell $(go_code) > $@)

# Compare git tag with current tag
# If git tag is newer than existing tag, remove git_describe_go and remake 
remove_older_describe_go : $(git_describe_go) force 
	@export data=$(shell $(go_code) | diff -q - $(git_describe_go) >/dev/null ; if [ $$? -ne 0 ] ; then (rm $(git_describe_go) ; printf "'new git tag: %s\n'" `$(git_cmd)`; $(MAKE) $(MAKEFLAGS) $(git_describe_go) > /dev/null ) ; fi) ; printf "$$data" ; ([ "$${#data}" -eq 0 ] || echo)

# git clean
.PHONY : clean
clean :
	go clean -i $(verbose_go_cmds) $(package)/...

.PHONY : force
force :
