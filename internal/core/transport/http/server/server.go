package core_http_server

import (
	"fmt"
	"errors"
	"context"
	"net/http"

	core_logger "github.com/kupr666/to-do-app/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux *http.ServeMux
	config Config
	log *core_logger.Logger
}

func NewHTTPServer(config Config, log *core_logger.Logger) *HTTPServer {
	return &HTTPServer{
		mux: http.NewServeMux(),
		config: config,
		log: log,
	}
}

func (h *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		h.mux.Handle(prefix+"/", http.StripPrefix(prefix, router))
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	server := &http.Server {
		Addr: h.config.Addr,
		Handler: h.mux,
	}

	ch := make(chan error, 1)

	h.log.Warn("Start http server", zap.String("addr", h.config.Addr))

	go func() {
		defer close(ch)

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <- ch:
		if err != nil {
			fmt.Errorf("Listen and serve HTTP: %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("shutdown HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			h.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		h.log.Warn("HTTP server stopped")
	}

	return nil	
}
