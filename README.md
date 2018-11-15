[![](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/aletheia7/gogitver) 

#### gogitver and [mage (go make)](https://github.com/magefile/mage)
A go package that inserts a git tag into your binary. Uses mage to efficiently 
compile your go application when the git tag changes. Other source code changes
will still trigger a compile. The git version/tag is created beforehand with:
```bash
git tag -a

# View the git tag: 
git describe --tags --long

# Instead of "go install" execute:
mage
```

#### Install and [mage](https://github.com/magefile/mage)
gogitver must be installed as a package under your application main
directory. E.g. if your application resides under ~/go/src/app, 
gogitver must be installed under ~/go/src/app/gogitver. ~/go/src/app/vendor/gogitver
is preferred.

```bash
# Example
mkdir ~/go/src/app/vendor/gogitver
cd ~/go/src/app
git clone --depth 1 https://github.com/aletheia7/gogitver.git vendor/gogitver
rm -fr vendor/gogitver/.git
# link mage_gogitver.go 
ln -s vendor/gogitver/mage_gogitver.go .
Add a mg.Dep(Gogitver) to your main magefile.go that should reside under ~/go/src/app .
```
Create a git tag:
```bash
git tag -a 1.0.7
```
Run mage:
```bash
mage
```

#### gogitver and [GNUMake](https://www.gnu.org/software/make/manual/make.html)
A go package that inserts a git tag into your binary. Uses make to efficiently 
compile your go application when the git tag changes. Other source code changes
will still trigger a compile. The git version/tag is created beforehand with:
```bash
git tag -a

# View the git tag: 
git describe --tags --long

# Instead of "go install" execute:
make
# Run "go install" and strip to remove debug symbols:
make strip
# Instead of "go clean" execute:
make clean
```

#### Install and [GNUMake](https://www.gnu.org/software/make/manual/make.html)
gogitver must be installed as a package under your application main
directory. E.g. if your application resides under ~/go/src/app, 
gogitver must be installed under ~/go/src/app/gogitver. ~/go/src/app/vendor/gogitver
is preferred.

```bash
# Example
mkdir ~/go/src/app/vendor/gogitver
cd ~/go/src/app
git clone --depth 1 https://github.com/aletheia7/gogitver.git vendor/gogitver
rm -fr vendor/gogitver/.git
# link makefile
ln -s vendor/gogitver/makefile .
```
Create a git tag:
```bash
git tag -a 1.0.7
```
Run make:
```bash
make
```

[GNUMake](https://www.gnu.org/software/make/manual/make.html)

Software Version Guidelines: [Semantic Versioning](http://semver.org)

Compare & parse semver version strings: [godoc go-semver](http://godoc.org/code.google.com/p/go-semver/version), [code: go-semver](https://code.google.com/p/go-semver/) 
#### Example

```go
package main
import (
	"fmt"
	// Do not use "github.com/alethiea7/gogitver"
	// It will not work
	"app/gogitver"
)

func main() {

	fmt.Println(gogitver.Git())	
	fmt.Println(gogitver.Tag())
}
```
##### Ouput
```bash
1.0.7-0-g8e9e07b
1.0.7
```

#### License 

Use of this source code is governed by a BSD-2-Clause license that can be found
in the LICENSE file.

[![BSD-2-Clause License](osi_logo_100X133_90ppi_0.png)](https://opensource.org/)
