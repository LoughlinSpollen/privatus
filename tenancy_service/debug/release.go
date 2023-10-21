//go:build !debug
// +build !debug

package debug

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var buildTags = []string{}

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)
	buildTags = append(buildTags, "release")
}

func Tags() []string {
	return buildTags
}
