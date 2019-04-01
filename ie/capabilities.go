package ie

import (
	"github.com/tebeka/selenium"
)

// CapabilitiesKey ...

const (
	CapabilitiesKey            = "se:ieOptions"
	UsePerProcessProxyKey      = "ie.usePerProcessProxy"
	IgnorProtectedModeSettings = "ignoreProtectedModeSettings"
	EnsureCleanSession         = "ie.ensureCleanSession"
	IgnoreZoomLevel            = "ignoreZoomSetting"
)

// Capabilities ...
type Capabilities struct {
	_opts map[string]interface{} // bool `json:"ie.usePerProcessProxy,omitempty"`
}

// NewCapabilities ...
func NewCapabilities() *Capabilities {
	return &Capabilities{_opts: map[string]interface{}{}}
}

// SetOption ...
func (c Capabilities) SetOption(key string, value interface{}) {
	c._opts[key] = value
}

// AddCap ...
func (c Capabilities) AddCap(cap selenium.Capabilities) {
	cap[CapabilitiesKey] = c._opts
}
