package main

import (
	"fmt"

	"gitlab.alibaba-inc.com/redcoast/version"
)

func main() {
	version.Parse()

	fmt.Printf("Version: %v\n", version.Version())
	fmt.Printf("The Last CommitID: %v\n", version.LastCommitID())
}
