//go:build wireinject
// +build wireinject

package internal

import "github.com/google/wire"

func New(config Config) App {
	wire.Build(NewApp, NewServer, NewRouter, NewCaptureCtl)

	return App{}
}
