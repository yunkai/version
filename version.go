package version

import (
	"flag"
	"fmt"
	"os"
)

// Version info set by user.
var Version string

// This function should be called in
// user's main function.
func Parse() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if vflag {
		fmt.Println(Version)
		os.Exit(0)
	}
}

var vflag bool

func init() {
	flag.BoolVar(&vflag, "version", false, "show version and exit")
}
