package version

import (
	"flag"
	"fmt"
	"os"
	"path"
)

var (
	vflag   bool
	version string
	commit  string
)

// This function should be called in
// user's main function.
func Parse() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if vflag {
		prg := path.Base(os.Args[0])
		fmt.Println(prg, version, commit)
		os.Exit(0)
	}
}

// Get project version.
func Version() string {
	return version
}

// Get the last commit id of source code.
func LastCommitID() string {
	return commit
}

func init() {
	flag.BoolVar(&vflag, "version", false, "show version and exit")
}
