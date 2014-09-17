#### gogitver 
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

#### Install 
gogitver must be installed as a package under your application main
directory. E.g. if your application resides under ~/go/src/app, 
gogitver must be installed under ~/go/src/app/gogitver.

```bash
# Example
mkdir ~/go/src/app/gogitver
cd ~/go/src/app
git clone --depth 1 https://github.com/aletheia7/gogitver.git
rm -fr gogitver/.git
# Copy makefile
cp gogitver/makefile .
```
Create a git tag:
```bash
git tag -a 1.0.7
```
Run make:
```bash
make
```
#### Documentation
[godoc gogitver](http://godoc.org/github.com/aletheia7/gogitver)

[GNUMake](https://www.gnu.org/software/make/manual/make.html)

Software Version Guidelines: [Semantic Versioning](http://semver.org)
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

![LGPL](http://www.gnu.org/graphics/lgplv3-147x51.png)
