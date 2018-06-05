package config

import (
	"github.com/spf13/afero"
	"github.com/spf13/pflag"

	"github.com/dcso/spyre"
	"github.com/dcso/spyre/log"
	"github.com/dcso/spyre/report"
)

var (
	Paths       []string
	MaxFileSize int64
	Hostname    string
)

// Fs is the "filesystem" in which configuration and rules are found.
// This can be provided through a ZIP file appended to the binary.
var Fs afero.Fs

func Init() error {
	pflag.StringSliceVarP(&Paths, "path", "p", defaultPaths,
		"paths to be scanned")
	pflag.Int64VarP(&MaxFileSize, "max-file-size", "", 32*1024*1024,
		"maximum size of individual files to be scanned, turn off by setting to 0 or negative value")
	pflag.StringVar(&spyre.Hostname, "set-hostname", spyre.DefaultHostname, "hostname")
	pflag.VarP(&log.GlobalLevel, "loglevel", "l", "loglevel")
	pflag.VarP(&report.Targets, "report", "r", "report target(s)")
	pflag.Parse()
	return nil
}
