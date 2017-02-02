# Version
一个小巧的版本管理库，可统一地管理Golang程序的`-version`选项，以及获取源码中最新的CommitID。

本项目的设计初衷是为了解决一个运维难题：**线上有很多程序，在出现Bug时，开发人员不清楚它的源码版本是哪个**。

## How To Use
### Step1
编译Golang程序时，创建Makefile，通过`go build`命令的`-ldflags`选项把程序的`版本号`和最新的`CommitID`编译到程序中：

```go
Version := 1.0.0
LastCommitID := $(shell git rev-parse HEAD)

all:
	go build -ldflags "\
		-X gitlab.alibaba-inc.com/redcoast/version.version=${Version} \
		-X gitlab.alibaba-inc.com/redcoast/version.commit=${LastCommitID} "\
		-o a.out
```

### Step2
在Golang源码中import本Project，并在main函数中调用`Parse()`函数，然后就可以通过`Version()`和`LastCommitID()`函数访问版本号和最新的CommitID：

```go
package main

import (
	"fmt"

	"github.com/yunkai/version"
)

func main() {
	version.Parse()

	fmt.Printf("Version: %v\n", version.Version())
	fmt.Printf("The Last CommitID: %v\n", version.LastCommitID())
}
```

以下是示例程序的执行结果：
```
$ ./a.out
Version: 1.0.0
The Last CommitID: 2dd3db1d85e7eade5c0b6309ce9c67a5312e9fde
```

还可以通过命令行选项`-version`查看到版本号和LastCommitID：
```bash
$ ./a.out -version
a.out 1.0.0 2dd3db1d85e7eade5c0b6309ce9c67a5312e9fde

```
