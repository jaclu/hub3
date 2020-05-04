package ikuzo

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/delving/hub3/config"
	"github.com/delving/hub3/ikuzo/logger"
	"github.com/delving/hub3/ikuzo/service/organization"
	"github.com/delving/hub3/ikuzo/service/x/bulk"
	"github.com/delving/hub3/ikuzo/service/x/ead"
	"github.com/delving/hub3/ikuzo/service/x/revision"
	"github.com/delving/hub3/ikuzo/storage/x/elasticsearch"
	"github.com/go-chi/chi"
)

// RouterFunc is a callback that registers routes to the ikuzo.Server.
type RouterFunc func(router chi.Router)

// Option is a closure to configure the Server.
// It is used in NewServer.
type Option func(*server) error

// SetPort sets the TCP port for the Server.
//
// The Server listens on :3000 by default.
func SetPort(port int) Option {
	return func(s *server) error {
		s.port = port
		return nil
	}
}

// SetMetricsPort sets the TCP port for the metrics server.
//
// No default. When set to 0 the metrics server is not started
func SetMetricsPort(port int) Option {
	return func(s *server) error {
		s.metricsPort = port
		return nil
	}
}

// SetLogger configures the global logger for the server.
func SetLogger(l *logger.CustomLogger) Option {
	return func(s *server) error {
		s.logger = l
		return nil
	}
}

// SetDisableRequestLogger disables logging of HTTP request
func SetDisableRequestLogger() Option {
	return func(s *server) error {
		s.disableRequestLogger = true
		return nil
	}
}

// SetMiddleware configures the global middleware for the HTTP router.
func SetMiddleware(middleware ...func(next http.Handler) http.Handler) Option {
	return func(s *server) error {
		s.middleware = append(s.middleware, middleware...)
		return nil
	}
}

// SetRouters adds all HTTP routes for the server.
func SetRouters(rb ...RouterFunc) Option {
	return func(s *server) error {
		s.routerFuncs = append(s.routerFuncs, rb...)
		return nil
	}
}

// SetOrganisationService configures the organization service.
// When no service is set a default transient memory-based service is used.
func SetOrganisationService(service *organization.Service) Option {
	return func(s *server) error {
		s.organizations = service
		return nil
	}
}

// SetRevisionService configures the organization service.
// When no service is set a default transient memory-based service is used.
func SetRevisionService(service *revision.Service) Option {
	return func(s *server) error {
		s.revision = service
		s.routerFuncs = append(s.routerFuncs, func(r chi.Router) {
			r.HandleFunc("/git/{user}/{collection}.git/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				p := strings.TrimPrefix(r.URL.Path, "/git")
				if !service.BareRepo {
					p = strings.ReplaceAll(p, ".git/", "/.git/")
				}
				r2 := new(http.Request)
				*r2 = *r
				r2.URL = new(url.URL)
				*r2.URL = *r.URL
				r2.URL.Path = p
				service.ServeHTTP(w, r2)
			}))
		})

		return nil
	}
}

func SetElasticSearchProxy(proxy *elasticsearch.Proxy) Option {
	return func(s *server) error {
		s.routerFuncs = append(s.routerFuncs,
			func(r chi.Router) {
				r.Handle("/{index}/_search", proxy)
				r.Handle("/{index}/{documentType}/_search", proxy)
			},
		)

		return nil
	}
}

func SetBuildVersionInfo(info *BuildVersionInfo) Option {
	return func(s *server) error {
		s.routerFuncs = append(s.routerFuncs,
			func(r chi.Router) {
				r.Get("/version", func(w http.ResponseWriter, r *http.Request) {
					s.respond(w, r, info, http.StatusOK)
				})
			},
		)

		return nil
	}
}

func SetEnableLegacyConfig() Option {
	return func(s *server) error {
		// this initializes the hub3 configuration object that has global state
		// TODO(kiivihal): remove this after legacy hub3/server/http/handlers are migrated
		config.InitConfig()

		return nil
	}
}

func SetEADService(svc *ead.Service) Option {
	return func(s *server) error {
		s.routerFuncs = append(s.routerFuncs,
			func(r chi.Router) {
				r.Post("/api/ead", svc.Upload)
			},
		)

		return nil
	}
}

func SetBulkService(svc *bulk.Service) Option {
	return func(s *server) error {
		s.routerFuncs = append(s.routerFuncs,
			func(r chi.Router) {
				r.Post("/api/index/bulk", svc.Handle)
			},
		)

		return nil
	}
}
