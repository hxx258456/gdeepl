/*
Copyright 2022 github.com/hxx258456/gdeepl
*/
package metadata

var Version = "1.0.0"

func GetVersion() string {
	if Version == "" {
		panic("Version is not set for fabric-ca library")
	}
	return Version
}
