package configuration

import (
	"github.com/mersanuzun/go-bff-boilerplate/lib/configuration/providers"
	"time"
)

type stringParamBuilder struct {
	f func() (string, error)
	paramBuilder
}

func newStringParamBuilder(provider providers.Provider, paramName string) *stringParamBuilder {
	var builder stringParamBuilder

	builder.provider = provider
	builder.param(paramName)

	return &builder
}

func (t *stringParamBuilder) param(paramName string) *stringParamBuilder {
	t.f = func() (string, error) {
		v, err := t.provider.GetString(paramName)
		return v, err
	}

	return t
}

func (t *stringParamBuilder) Build() func() string {
	return func() string {
		value, err := t.f()

		if err != nil {
			panic(err)
		} else {
			t.fetchedTime = time.Now()
		}

		return value
	}
}
