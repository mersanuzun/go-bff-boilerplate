package configuration

import (
	"github.com/mersanuzun/go-bff-boilerplate/lib/configuration/providers"
	"time"
)

type paramBuilder struct {
	fetchedTime time.Time
	provider    providers.Provider
}
