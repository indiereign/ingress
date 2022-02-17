module github.com/caddyserver/ingress

go 1.16

replace github.com/indiereign/shift72-ingress-config => ../shift72-caddy-config

require (
	github.com/caddyserver/caddy/v2 v2.4.6
	github.com/caddyserver/certmagic v0.15.2
	github.com/google/uuid v1.3.0
	github.com/indiereign/shift72-ingress-config v0.0.0-00010101000000-000000000000
	github.com/mitchellh/mapstructure v1.1.2
	github.com/pires/go-proxyproto v0.3.1
	github.com/pkg/errors v0.9.1
	go.uber.org/zap v1.19.0
	gopkg.in/go-playground/pool.v3 v3.1.1
	k8s.io/api v0.19.4
	k8s.io/apimachinery v0.19.4
	k8s.io/client-go v0.19.4
)

// update
replace (
	k8s.io/api => k8s.io/api v0.19.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.4
	k8s.io/client-go => k8s.io/client-go v0.19.4
)
