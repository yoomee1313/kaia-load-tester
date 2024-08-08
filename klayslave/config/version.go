package config

import (
	"fmt"
)

// sets build options from ldflags.
const (
	Major = 1
	Minor = 0
	Patch = 1
)

var (
	Commit    = ""
	Branch    = ""
	Tag       = ""
	BuildDate = ""
	BuildUser = ""
	Version   = func() string {
		v := fmt.Sprintf("v%d.%d.%d", Major, Minor, Patch)
		return v
	}()
)

func GetVersionWithCommit() string {
	if len(Commit) >= 8 {
		return Version + "-" + Commit[:8]
	}
	return Version
}

func String() string {
	version := Version
	if len(Commit) >= 7 {
		version += "-" + Commit[:7]
	}
	if Tag != "" && Tag != "undefined" {
		version = Tag
	}
	fmt.Printf("Version :\t%s\n", version)
	fmt.Printf("git.Branch :\t%s\n", Branch)
	fmt.Printf("git.Commit :\t%s\n", Commit)
	fmt.Printf("git.Tag :\t%s\n", Tag)
	fmt.Printf("build.Date :\t%s\n", BuildDate)
	fmt.Printf("build.User :\t%s\n", BuildUser)
	return version
}
