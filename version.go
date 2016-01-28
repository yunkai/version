package version

import (
	"flag"
	"fmt"
	"os"
	"path"
)

// Version info set by user.
var Version string

// The last commit id of user's project.
var LastCommitID string

// This function should be called in
// user's main function.
func Parse() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if vflag {
		prg := path.Base(os.Args[0])
		fmt.Println(prg, Version, LastCommitID)
		os.Exit(0)
	}
}

var vflag bool

func init() {
	flag.BoolVar(&vflag, "version", false, "show version and exit")
}
