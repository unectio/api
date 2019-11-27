package api

import (
	"github.com/unectio/util"
)

type SpecEntry struct {
	Type		string			`yaml:"type"`
	Spec		util.YAMLRaw		`yaml:"spec"`
}
