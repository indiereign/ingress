package config

import (
	"github.com/caddyserver/caddy/v2"
)

// StorageValues represents the config for certmagic storage providers.
type StorageValues struct {
	Namespace string `json:"namespace"`
	LeaseId   string `json:"leaseId"`
}

// Storage represents the certmagic storage configuration.
type Storage struct {
	System string `json:"module"`
	StorageValues
}

// Config represents a caddy2 config file.
type Config struct {
	Admin   caddy.AdminConfig      `json:"admin,omitempty"`
	Storage Storage                `json:"storage"`
	Apps    map[string]interface{} `json:"apps"`
	Logging caddy.Logging          `json:"logging"`
}

type Converter interface {
	ConvertToCaddyConfig(namespace string, store *Store) (interface{}, error)
}
