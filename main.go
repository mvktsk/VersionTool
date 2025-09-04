package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"versiontool/version"
)

var (
    appVersion  string = "0.0.0"
)


func main() {
	fmt.Println("Version:", appVersion)
	versionPart := flag.String("version-part", "minor", "increment type (patch, minor, major)")
	flag.Parse()
	slog.Info("Start application")
	currentVer, err := version.Parse(appVersion)
	if err != nil {
		slog.Error("Error parsing version", "Error", err.Error())
	}
	currentVer.Increment(*versionPart)
	os.Exit(0)
}