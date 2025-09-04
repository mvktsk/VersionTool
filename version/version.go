package version

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"
)

const (
	semVerFormat = "%d.%d.%d"
)

// Version represents a version in SemVer or CalVer format
type Version struct {
	Major, Minor, Patch int
	Build, Prerelease   string
}

func (v *Version) IncrementPatch() {
	v.Build = ""
	v.Patch++
}

func (v *Version) IncrementMinor() {
	v.Build = ""
	v.Minor++
	v.Patch = 0
}

func (v *Version) IncrementMajor() {
	v.Major++
	v.Minor = 0
	v.Patch = 0
	v.Build = ""
}

// ParseVersion parses a version string into a Version struct
func Parse(versionStr string) (*Version, error) {

	// Split build metadata if present
	parts := strings.Split(versionStr, "+")
	version := parts[0]
	var build string
	if len(parts) > 1 {
		if parts[1] == "" {
			return nil, errors.New("Empty build metadata is not allowed")
		}
		build = parts[1]
	}

	// Parse version numbers
	if len(strings.Split(version, ".")) != 3 {
		return nil, errors.New("invalid version format: " + versionStr + " (expected format X.Y.Z)")
	}

	// Parse version numbers
	var v Version
	_, err := fmt.Sscanf(version, "%d.%d.%d", &v.Major, &v.Minor, &v.Patch)
	if err != nil {
		slog.Error("Error parsing version", "Error", err.Error())
		return nil, err
	}
	
	// Set build metadata
	v.Build = build

	return &v, nil
}


func (v *Version) Increment(incrementType string) (error) {
	oldVersion := v.String()
	
	switch incrementType {
	case "patch":
		v.IncrementPatch()
	case "minor":
		v.IncrementMinor()
	case "major":
		v.IncrementMajor()
	default:
		return errors.New("invalid increment type: " + incrementType + " (expected patch, minor, major, or calendar)", ) 
	}
	
	newVersion := v.String()
	slog.Info("Version successfully updated", "Old version", oldVersion, "New version",  newVersion)
		return nil
}


// String returns the string representation of the version
func (v *Version) String() string {
	format := semVerFormat

	version := fmt.Sprintf(format, v.Major, v.Minor, v.Patch)
	if v.Build != "" {
		version += "+" + v.Build
	}
	return version
}

