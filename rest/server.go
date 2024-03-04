package rest

import (
	"net/http"

	"github.com/rmukhamet/finance/config"
)

var (
	healthy int32
	ready   int32
	// watcher *fscache.Watcher
)

type Logger interface {
	Info(...interface{})
}

func NewServer(
	logger Logger,
	config *config.Config,
	// commentStore *commentStore,
	// anotherStore *anotherStore.
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		logger,
		// config,
		// commentStore,
		// anotherStore,
	)
	var handler http.Handler = mux
	// handler = someMiddleware(handler)
	// handler = someMiddleware2(handler)
	// handler = someMiddleware3(handler)
	return handler
}
