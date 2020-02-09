package config

import (
	"github.com/mersanuzun/go-bff-boilerplate/lib/configuration"
	. "github.com/mersanuzun/go-bff-boilerplate/lib/configuration/providers"
)

// App
var Env = configuration.From(Environment()).StringParam("ENV").DefaultTo("local").Build()
var AppName = configuration.From(Environment()).StringParam("APP_NAME").DefaultTo("Go BFF Boilerplate With Echo").Build()
var Version = configuration.From(Environment()).StringParam("VERSION").DefaultTo("release").Build()
var Port = configuration.From(Environment()).StringParam("SERVER_PORT").DefaultTo("3000").Build()


var JsonPlaceHolderApi = configuration.
	From(Environment()).
	StringParam("JSON_PLACE_HOLDER_API_URL").
	CacheOnce().
	DefaultTo("https://jsonplaceholder.typicode.com").
	Build()
