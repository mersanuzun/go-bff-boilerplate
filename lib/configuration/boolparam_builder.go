package configuration

import (
	"github.com/mersanuzun/go-bff-boilerplate/lib/configuration/providers"
	"time"
)

type boolParamBuilder struct {
	f func() (bool, error)
	paramBuilder
}

func newBoolParamBuilder(provider providers.Provider, paramName string) *boolParamBuilder {
	var builder boolParamBuilder

	builder.provider = provider
	builder.param(paramName)

	return &builder
}

func (t *boolParamBuilder) param(paramName string) *boolParamBuilder {
	t.f = func() (bool, error) {
		v, err := t.provider.GetBool(paramName)
		return v, err
	}

	return t
}

func (t *boolParamBuilder) Build() func() bool {
	return func() bool {
		value, err := t.f()

		if err != nil {
			panic(err)
		} else {
			t.fetchedTime = time.Now()
		}

		return value
	}
}
