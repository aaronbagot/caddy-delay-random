// Copyright 2020 Patrick Easters
// SPDX-License-Identifier: Apache-2.0
// Modified by aaronbagot, 2026

package delay_random

import (
	"fmt"
	"strconv"
	"time"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("delay_random", parseCaddyfileHandler)
}

// parseCaddyfileHandler unmarshals tokens from h into a new middleware handler value.
func parseCaddyfileHandler(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	m := new(Middleware)
	if err := m.UnmarshalCaddyfile(h.Dispenser); err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
//
// Syntax:
//
//	delay_random <min_delay> <max_delay> [<probability>]
//
// Examples:
//     delay_random 500ms 2s
//     delay_random 300ms 1.5s 0.8
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	d.Next() // skip delay_random

	var minStr, maxStr, probStr string
	if !d.Args(&minStr, &maxStr) {
		return d.ArgErr()
	}
	if d.NextArg() {
		probStr = d.Val()
	}
	if d.NextArg() {
		return d.ArgErr()
	}

	var err error

	m.MinDelay, err = time.ParseDuration(minStr)
	if err != nil {
		return fmt.Errorf("failed to parse min_delay: %w", err)
	}

	m.MaxDelay, err = time.ParseDuration(maxStr)
	if err != nil {
		return fmt.Errorf("failed to parse max_delay: %w", err)
	}

	if probStr != "" {
		var prob float64
		prob, err = strconv.ParseFloat(probStr, 64)
		if err != nil {
			return fmt.Errorf("failed to parse probability: %w", err)
		}
		m.Probability = &prob
	}
	// If no probability given, Provision() will set default to 1.0

	return nil
}
