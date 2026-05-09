# caddy-delay-random
[![Go Reference](https://pkg.go.dev/badge/github.com/aaronbagot/caddy-delay-random.svg)](https://pkg.go.dev/github.com/aaronbagot/caddy-delay-random)

A Caddy HTTP handler to introduce a random delay for some requests.

## Purpose
This module helps reproduce issues caused by slow HTTP requests and is useful for chaos testing or simulating slow backend services.

## Usage (Caddyfile)
```caddyfile
delay_random <min_delay> <max_delay> [<probability>]
```

* `min_delay` — the minimum delay duration (e.g. `100ms`, `500ms`, `1s`)
* `max_delay` — the maximum delay duration (e.g. `2s`, `5s`, `10s`)
* `probability` (optional) — Float between `0.0` and `1.0`. Defaults to `1.0` (100%).

Durations follow [Go's duration format](https://pkg.go.dev/time#ParseDuration) (e.g. `500ms`, `2s`, `1m30s`).

# Examples

## Always delay (100%) between 500ms and 1.5s
```markdown
#### `Caddyfile`
```caddyfile
{
    order delay_random before reverse_proxy
}

http://localhost:8000 {
    delay_random 500ms 1.5s
    reverse_proxy 127.0.0.1:9000
}
```
```

## Delay between 1s and 5s for 30% of requests
```markdown
#### `Caddyfile`
```caddyfile
{
    order delay_random before reverse_proxy
}

http://localhost:8000 {
    delay_random 1s 5s 0.3
    reverse_proxy 127.0.0.1:9000
}
```
```

## JSON Configuration
```markdown
#### `caddy.json`
```json
{
    "handler": "delay_random",
    "min_delay": "500ms",
    "max_delay": "2s",
    "probability": 0.75
}
```
```

# License
This project is a fork of [caddy-random-delay](https://github.com/patrickeasters/caddy-random-delay) by Patrick Easters, licensed under the Apache License 2.0.  
Modifications copyright © 2026 Aaron Bagot.
