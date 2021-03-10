// Code generated by "mapstructure-to-hcl2 -type ShutdownConfig"; DO NOT EDIT.

package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatShutdownConfig is an auto-generated flat version of ShutdownConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatShutdownConfig struct {
	Command         *string `mapstructure:"shutdown_command" cty:"shutdown_command" hcl:"shutdown_command"`
	Timeout         *string `mapstructure:"shutdown_timeout" cty:"shutdown_timeout" hcl:"shutdown_timeout"`
	DisableShutdown *bool   `mapstructure:"disable_shutdown" cty:"disable_shutdown" hcl:"disable_shutdown"`
}

// FlatMapstructure returns a new FlatShutdownConfig.
// FlatShutdownConfig is an auto-generated flat version of ShutdownConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*ShutdownConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatShutdownConfig)
}

// HCL2Spec returns the hcl spec of a ShutdownConfig.
// This spec is used by HCL to read the fields of ShutdownConfig.
// The decoded values from this spec will then be applied to a FlatShutdownConfig.
func (*FlatShutdownConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"shutdown_command": &hcldec.AttrSpec{Name: "shutdown_command", Type: cty.String, Required: false},
		"shutdown_timeout": &hcldec.AttrSpec{Name: "shutdown_timeout", Type: cty.String, Required: false},
		"disable_shutdown": &hcldec.AttrSpec{Name: "disable_shutdown", Type: cty.Bool, Required: false},
	}
	return s
}
