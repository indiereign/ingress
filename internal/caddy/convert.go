package caddy

import (
	"encoding/json"

	caddy "github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/caddyserver/caddy/v2/modules/caddytls"
	"github.com/caddyserver/ingress/pkg/config"
)

type Converter struct{}

const (
	HttpServer    = "ingress_server"
	MetricsServer = "metrics_server"
)

func metricsServer(enabled bool) *caddyhttp.Server {
	handler := json.RawMessage(`{ "handler": "static_response" }`)
	if enabled {
		handler = json.RawMessage(`{ "handler": "metrics" }`)
	}

	return &caddyhttp.Server{
		Listen:    []string{":9765"},
		AutoHTTPS: &caddyhttp.AutoHTTPSConfig{Disabled: true},
		Routes: []caddyhttp.Route{{
			HandlersRaw: []json.RawMessage{handler},
			MatcherSetsRaw: []caddy.ModuleMap{{
				"path": caddyconfig.JSON(caddyhttp.MatchPath{"/metrics"}, nil),
			}},
		}},
	}
}

func newConfig(namespace string, store *config.Store) (*config.Config, error) {
	cfg := &config.Config{
		Logging: caddy.Logging{},
		Apps: map[string]interface{}{
			"tls": &caddytls.TLS{
				CertificatesRaw: caddy.ModuleMap{},
			},
			"http": &caddyhttp.App{
				Servers: map[string]*caddyhttp.Server{
					MetricsServer: metricsServer(store.ConfigMap.Metrics),
					HttpServer: {
						AutoHTTPS: &caddyhttp.AutoHTTPSConfig{},
						// Listen to both :80 and :443 ports in order
						// to use the same listener wrappers (PROXY protocol use it)
						Listen: []string{":80", ":443"},
					},
				},
			},
		},
		Storage: config.Storage{
			System: "secret_store",
			StorageValues: config.StorageValues{
				Namespace: namespace,
				LeaseId:   store.Options.LeaseId,
			},
		},
	}

	return cfg, nil
}

func (c Converter) ConvertToCaddyConfig(namespace string, store *config.Store) (interface{}, error) {
	cfg, err := newConfig(namespace, store)
	if err != nil {
		return cfg, err
	}

	err = LoadIngressConfig(cfg, store)
	if err != nil {
		return cfg, err
	}

	err = LoadConfigMapOptions(cfg, store)
	if err != nil {
		return cfg, err
	}

	err = LoadTLSConfig(cfg, store)
	if err != nil {
		return cfg, err
	}

	return cfg, err
}
