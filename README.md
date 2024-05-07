# caddy-random-delay
[![Go Reference](https://pkg.go.dev/badge/github.com/patrickeasters/caddy-random-delay.svg)](https://pkg.go.dev/github.com/patrickeasters/caddy-random-delay)

A caddy handler to introduce a delay for some requests

# Purpose
This module was written to facilitate reproducing a problem that occurred when some HTTP requests began taking too long, resulting in client-side timeouts. While one usually doesn't want to make requests take longer, this module could be useful for chaos testing or simulating slow HTTP servers.

# Usage
```
random_delay <percent_delayed> <duration>
```

The module has only two options:
* `percent_delayed` is a float that determines the chance a request will be delayed (e.g. `0.2`, `0.69`, `1.0`)
* `duration` is a string that determines the duration of the delay injected by the module following [Go's duration format](https://pkg.go.dev/time#ParseDuration) (e.g. `10s`, `5m`, `200ms`)

A full Caddyfile example is below. This would introduce a 10 second delay to approximately 50% of requests.

```
{
	order random_delay before reverse_proxy
}

http://localhost:8000 {
    random_delay 0.5 10s
	reverse_proxy https://icanhazip.com {
		header_up Host {upstream_hostport}
	}
	log {
		format console
	}
}
```
