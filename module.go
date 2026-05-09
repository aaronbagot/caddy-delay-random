// Copyright 2020 Patrick Easters
// SPDX-License-Identifier: Apache-2.0
// Modified by aaronbagot, 2026

package delay_random

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(Middleware{})
}

// Middleware implements an HTTP handler that delays requests by a random duration.
type Middleware struct {
	MinDelay     time.Duration  `json:"min_delay"`
	MaxDelay     time.Duration  `json:"max_delay"`
	Probability  *float64       `json:"probability,omitempty"`

	randomSource *rand.Rand
}

// CaddyModule returns the Caddy module information.
func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.delay_random",
		New: func() caddy.Module { return new(Middleware) },
	}
}

// Provision implements caddy.Provisioner
func (m *Middleware) Provision(ctx caddy.Context) error {
	// If probability is not set (nil), default to 1.0 (100%)
	if m.Probability == nil {
		defaultProb := 1.0
		m.Probability = &defaultProb
	}

	src := rand.NewSource(time.Now().UnixNano())
	m.randomSource = rand.New(src)
	return nil
}

// Validate implements caddy.Validator.
func (m *Middleware) Validate() error {
	if m.MinDelay < 0 {
		return fmt.Errorf("min_delay cannot be negative")
	}
	if m.MaxDelay < m.MinDelay {
		return fmt.Errorf("max_delay must be greater than or equal to min_delay")
	}
	if m.Probability == nil {
		return fmt.Errorf("internal error: probability is nil")
	}
	if *m.Probability < 0 || *m.Probability > 1 {
		return fmt.Errorf("probability must be between 0 and 1")
	}
	return nil
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Middleware)(nil)
	_ caddy.Validator             = (*Middleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
	_ caddyfile.Unmarshaler       = (*Middleware)(nil)
)