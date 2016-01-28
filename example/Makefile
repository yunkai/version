Version := 1.0.0
LastCommitID := $(shell git rev-parse HEAD)

all:
	go build -ldflags "\
		-X gitlab.alibaba-inc.com/redcoast/version.Version=${Version} \
		-X gitlab.alibaba-inc.com/redcoast/version.LastCommitID=${LastCommitID} "\
		-o a.out
