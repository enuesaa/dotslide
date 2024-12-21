//go:build wireinject
// +build wireinject

package internal

import "github.com/google/wire"

func New() App {
	wire.Build(NewApp)

	return App{}
}
