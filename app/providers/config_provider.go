package providers

import (
	"github.com/raiworks/rapidgo/v2/core/config"
	"github.com/raiworks/rapidgo/v2/core/container"
)

// ConfigProvider loads environment configuration.
// Must be registered first ΓÇö all other providers may read config
// values during their own Register() calls.
type ConfigProvider struct{}

// Register loads the .env file via config.Load().
func (p *ConfigProvider) Register(c *container.Container) {
	config.Load()
}

// Boot is a no-op. Config is fully available after Load().
func (p *ConfigProvider) Boot(c *container.Container) {}