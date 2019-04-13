package main

import (
	"path/filepath"
	"strings"
)

// taken from Athens ;)
func moduleVersionFromPath(path string) (string, string) {
	segments := strings.Split(path, "/@v/")
	if len(segments) != 2 {
		return "", ""
	}
	version := strings.TrimSuffix(segments[1], filepath.Ext(segments[1]))
	return segments[0], version
}
