//go:build wasm
// +build wasm

package app

import (
	"runtime"

	"github.com/di4f/pwa/errors"
)

const (
	appJS        = ""
	appWorkerJS  = ""
	manifestJSON = ""
	appCSS       = ""
)

var (
	errBadInstruction = errors.New("unsupported instruction").
		WithTag("architecture", runtime.GOARCH)
)

func GenerateStaticWebsite(dir string, h *Handler, pages ...string) error {
	panic(errBadInstruction)
}

func wasmExecJS() string {
	return ""
}
