// Package flagbytes implements a flag for specifying byte sizes in a human
// format.
//
// For details about the format, look at the documentation for the underlying
// library here:
//   http://go.pkgdoc.org/github.com/dustin/go-humanize
package flagbytes

import (
	"flag"
	"github.com/dustin/go-humanize"
)

type bytes struct {
	dest *uint64
}

func BytesVar(dest *uint64, name string, value string, usage string) {
	var err error
	*dest, err = humanize.ParseBytes(value)
	if err != nil {
		panic(err)
	}
	flag.Var(&bytes{dest}, name, usage)
}

func Bytes(name string, value string, usage string) *uint64 {
	d := new(uint64)
	BytesVar(d, name, value, usage)
	return d
}

func (d *bytes) String() string {
	return humanize.Bytes(*d.dest)
}

func (d *bytes) Set(val string) (err error) {
	*d.dest, err = humanize.ParseBytes(val)
	return
}
