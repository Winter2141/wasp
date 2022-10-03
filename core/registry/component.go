// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package registry

import (
	"go.uber.org/dig"

	"github.com/iotaledger/hive.go/core/app"
	"github.com/iotaledger/wasp/core/database"
	"github.com/iotaledger/wasp/packages/registry"
)

func init() {
	CoreComponent = &app.CoreComponent{
		Component: &app.Component{
			Name:      "Registry",
			Params:    params,
			Provide:   provide,
			Configure: configure,
		},
	}
}

var (
	CoreComponent *app.CoreComponent

	defaultRegistry registry.Registry
)

func provide(c *dig.Container) error {
	if err := c.Provide(func() *registry.Config {
		return &registry.Config{UseText: ParamsRegistry.UseText, Filename: ParamsRegistry.FileName}
	}); err != nil {
		CoreComponent.LogPanic(err)
	}

	return nil
}

func configure() error {
	logger := CoreComponent.Logger()

	if ParamsRegistry.UseText {
		defaultRegistry = registry.NewTextRegistry(logger, ParamsRegistry.FileName)
	} else {
		defaultRegistry = registry.NewRegistry(logger, database.GetRegistryKVStore())
	}

	return nil
}

// DefaultRegistry returns an initialized default registry.
func DefaultRegistry() registry.Registry {
	return defaultRegistry
}
