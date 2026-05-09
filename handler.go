// Copyright 2020 Patrick Easters
// SPDX-License-Identifier: Apache-2.0
// Modified by aaronbagot, 2026

package delay_random

import (
	"net/http"
	"time"

	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// ServeHTTP implements caddyhttp.MiddlewareHandler
func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// Probability should not be nil or <0, just in case
	if m.Probability == nil || *m.Probability <= 0 {
		return next.ServeHTTP(w, r)
	}

	prob := *m.Probability
	if m.randomSource.Float64() < prob {
		delta := m.MaxDelay - m.MinDelay
		randomDelay := m.MinDelay
		if delta > 0 {
			randomDelay += time.Duration(m.randomSource.Int63n(int64(delta) + 1))
		}

		timer := time.NewTimer(randomDelay)
		select {
		case <-r.Context().Done():
			if !timer.Stop() {
				<-timer.C
			}
		case <-timer.C:
		}
	}

	return next.ServeHTTP(w, r)
}