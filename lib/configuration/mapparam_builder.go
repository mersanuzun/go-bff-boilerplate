package configuration

import (
	"github.com/mersanuzun/go-bff-boilerplate/lib/configuration/providers"
	"time"
)

type mapParamBuilder struct {
	f func() (map[string]string, error)
	paramBuilder
}

func newMapParamBuilder(provider providers.Provider, paramName string) *mapParamBuilder {
	var builder mapParamBuilder

	builder.provider = provider
	builder.param(paramName)

	return &builder
}

func (t *mapParamBuilder) param(paramName string) *mapParamBuilder {
	t.f = func() (map[string]string, error) {
		v, err := t.provider.GetMap(paramName)
		return v, err
	}

	return t
}

func (t *mapParamBuilder) Build() func() map[string]string {
	return func() map[string]string {
		value, err := t.f()
		if err != nil {
			panic(err)
		} else {
			t.fetchedTime = time.Now()
		}

		return value
	}
}
