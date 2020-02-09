package configuration

import (
	"github.com/mersanuzun/go-bff-boilerplate/lib/configuration/providers"
	"time"
)

type intParamBuilder struct {
	f func() (int, error)
	paramBuilder
}

func newIntParamBuilder(provider providers.Provider, paramName string) *intParamBuilder {
	var builder intParamBuilder

	builder.provider = provider
	builder.param(paramName)

	return &builder
}

func (t *intParamBuilder) param(paramName string) *intParamBuilder {
	t.f = func() (int, error) {
		v, err := t.provider.GetInt(paramName)
		return v, err
	}

	return t
}

func (t *intParamBuilder) Build() func() int {
	return func() int {
		value, err := t.f()

		if err != nil {
			panic(err)
		} else {
			t.fetchedTime = time.Now()
		}

		return value
	}
}
