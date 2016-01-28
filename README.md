# Version
A tiny library to manage `-version` flag and last commit id of source code.

## How To Use
在源码中import本Project，并在main函数中调用`version.Parse()`函数。

```go
package main

import (
	"fmt"

	"gitlab.alibaba-inc.com/redcoast/version"
)

func main() {
	version.Parse()

	// Your can access `Version` and `LastCommitID` variables
	// after `version.Parse()`.
	fmt.Printf("Version: %v\n", version.Version)
	fmt.Printf("The Last CommitID: %v\n", version.LastCommitID)
}
```

同时创建Makefile，通过`ldflags`选项把最新的CommitId值编译到程序中：

```go
Version := 1.0.0
LastCommitID := $(shell git rev-parse HEAD)

all:
	go build -ldflags "\
		-X gitlab.alibaba-inc.com/redcoast/version.Version=${Version} \
		-X gitlab.alibaba-inc.com/redcoast/version.LastCommitID=${LastCommitID} "\
		-o a.out
```

然后就可以通过命令行选项`-commitid`或在源码中直接访问`version.CommitID`获取源码中最新的Commit ID：

```bash
$ ./a.out -version
a.out 1.0 2dd3db1d85e7eade5c0b6309ce9c67a5312e9fde

$ ./a.out
Version: 1.0
The Last CommitID: 2dd3db1d85e7eade5c0b6309ce9c67a5312e9fde

```
