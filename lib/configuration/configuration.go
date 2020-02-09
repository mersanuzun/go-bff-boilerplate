package configuration

import "github.com/mersanuzun/go-bff-boilerplate/lib/configuration/providers"

func From(provider providers.Provider) *builderFactory {
	return &builderFactory{provider}
}

type builderFactory struct {
	provider providers.Provider
}

func (b *builderFactory) StringParam(paramName string) *stringParamBuilder {
	return newStringParamBuilder(b.provider, paramName)
}

func (b *builderFactory) IntParam(paramName string) *intParamBuilder {
	return newIntParamBuilder(b.provider, paramName)
}

func (b *builderFactory) BoolParam(paramName string) *boolParamBuilder {
	return newBoolParamBuilder(b.provider, paramName)
}

func (b *builderFactory) MapParam(paramName string) *mapParamBuilder {
	return newMapParamBuilder(b.provider, paramName)
}
