package app

import (
	"github.com/raiworks/rapidgo/v2/core/plugin"
	exampleplugin "github.com/raiworks/rapidgo-starter/plugins/example"
)

// RegisterPlugins registers all application plugins with the manager.
func RegisterPlugins(m *plugin.PluginManager) {
	m.Add(exampleplugin.New())
}