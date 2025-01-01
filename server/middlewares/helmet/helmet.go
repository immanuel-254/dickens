package helmet

import (
	"fmt"
	"net/http"
)

// Config defines the configuration for the middleware.
type Config struct {
	Next                      func(r *http.Request) bool
	XSSProtection             string
	ContentTypeNosniff        string
	XFrameOptions             string
	ContentSecurityPolicy     string
	CSPReportOnly             bool
	ReferrerPolicy            string
	PermissionPolicy          string
	CrossOriginEmbedderPolicy string
	CrossOriginOpenerPolicy   string
	CrossOriginResourcePolicy string
	OriginAgentCluster        string
	XDNSPrefetchControl       string
	XDownloadOptions          string
	XPermittedCrossDomain     string
	HSTSMaxAge                int
	HSTSExcludeSubdomains     bool
	HSTSPreloadEnabled        bool
}

// ConfigDefault provides default configuration values.
var ConfigDefault = Config{
	XSSProtection:             "0",
	ContentTypeNosniff:        "nosniff",
	XFrameOptions:             "SAMEORIGIN",
	ReferrerPolicy:            "no-referrer",
	CrossOriginEmbedderPolicy: "require-corp",
	CrossOriginOpenerPolicy:   "same-origin",
	CrossOriginResourcePolicy: "same-origin",
	OriginAgentCluster:        "?1",
	XDNSPrefetchControl:       "off",
	XDownloadOptions:          "noopen",
	XPermittedCrossDomain:     "none",
}

// New creates the middleware handler for `net/http`.
func New(config ...Config) func(next http.Handler) http.Handler {
	// Initialize the configuration.
	cfg := configDefault(config...)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if the middleware should skip the request.
			if cfg.Next != nil && cfg.Next(r) {
				next.ServeHTTP(w, r)
				return
			}

			// Set headers based on the configuration.
			if cfg.XSSProtection != "" {
				w.Header().Set("X-XSS-Protection", cfg.XSSProtection)
			}
			if cfg.ContentTypeNosniff != "" {
				w.Header().Set("X-Content-Type-Options", cfg.ContentTypeNosniff)
			}
			if cfg.XFrameOptions != "" {
				w.Header().Set("X-Frame-Options", cfg.XFrameOptions)
			}
			if cfg.CrossOriginEmbedderPolicy != "" {
				w.Header().Set("Cross-Origin-Embedder-Policy", cfg.CrossOriginEmbedderPolicy)
			}
			if cfg.CrossOriginOpenerPolicy != "" {
				w.Header().Set("Cross-Origin-Opener-Policy", cfg.CrossOriginOpenerPolicy)
			}
			if cfg.CrossOriginResourcePolicy != "" {
				w.Header().Set("Cross-Origin-Resource-Policy", cfg.CrossOriginResourcePolicy)
			}
			if cfg.OriginAgentCluster != "" {
				w.Header().Set("Origin-Agent-Cluster", cfg.OriginAgentCluster)
			}
			if cfg.ReferrerPolicy != "" {
				w.Header().Set("Referrer-Policy", cfg.ReferrerPolicy)
			}
			if cfg.XDNSPrefetchControl != "" {
				w.Header().Set("X-DNS-Prefetch-Control", cfg.XDNSPrefetchControl)
			}
			if cfg.XDownloadOptions != "" {
				w.Header().Set("X-Download-Options", cfg.XDownloadOptions)
			}
			if cfg.XPermittedCrossDomain != "" {
				w.Header().Set("X-Permitted-Cross-Domain-Policies", cfg.XPermittedCrossDomain)
			}

			// Handle HSTS headers.
			if r.TLS != nil && cfg.HSTSMaxAge > 0 {
				subdomains := ""
				if !cfg.HSTSExcludeSubdomains {
					subdomains = "; includeSubDomains"
				}
				if cfg.HSTSPreloadEnabled {
					subdomains += "; preload"
				}
				w.Header().Set("Strict-Transport-Security", fmt.Sprintf("max-age=%d%s", cfg.HSTSMaxAge, subdomains))
			}

			// Handle Content-Security-Policy headers.
			if cfg.ContentSecurityPolicy != "" {
				if cfg.CSPReportOnly {
					w.Header().Set("Content-Security-Policy-Report-Only", cfg.ContentSecurityPolicy)
				} else {
					w.Header().Set("Content-Security-Policy", cfg.ContentSecurityPolicy)
				}
			}

			// Handle Permissions-Policy headers.
			if cfg.PermissionPolicy != "" {
				w.Header().Set("Permissions-Policy", cfg.PermissionPolicy)
			}

			// Proceed to the next handler.
			next.ServeHTTP(w, r)
		})
	}
}

// Helper function to apply default configuration values.
func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	// Apply default values where not set.
	if cfg.XSSProtection == "" {
		cfg.XSSProtection = ConfigDefault.XSSProtection
	}
	if cfg.ContentTypeNosniff == "" {
		cfg.ContentTypeNosniff = ConfigDefault.ContentTypeNosniff
	}
	if cfg.XFrameOptions == "" {
		cfg.XFrameOptions = ConfigDefault.XFrameOptions
	}
	if cfg.ReferrerPolicy == "" {
		cfg.ReferrerPolicy = ConfigDefault.ReferrerPolicy
	}
	if cfg.CrossOriginEmbedderPolicy == "" {
		cfg.CrossOriginEmbedderPolicy = ConfigDefault.CrossOriginEmbedderPolicy
	}
	if cfg.CrossOriginOpenerPolicy == "" {
		cfg.CrossOriginOpenerPolicy = ConfigDefault.CrossOriginOpenerPolicy
	}
	if cfg.CrossOriginResourcePolicy == "" {
		cfg.CrossOriginResourcePolicy = ConfigDefault.CrossOriginResourcePolicy
	}
	if cfg.OriginAgentCluster == "" {
		cfg.OriginAgentCluster = ConfigDefault.OriginAgentCluster
	}
	if cfg.XDNSPrefetchControl == "" {
		cfg.XDNSPrefetchControl = ConfigDefault.XDNSPrefetchControl
	}
	if cfg.XDownloadOptions == "" {
		cfg.XDownloadOptions = ConfigDefault.XDownloadOptions
	}
	if cfg.XPermittedCrossDomain == "" {
		cfg.XPermittedCrossDomain = ConfigDefault.XPermittedCrossDomain
	}

	return cfg
}