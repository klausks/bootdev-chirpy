package admin

import (
	"fmt"
	"net/http"
)

type MetricsHandler struct {
	metrics metrics
}

func NewHandler() *MetricsHandler {
	return &MetricsHandler{metrics: metrics{}}
}

func (handler *MetricsHandler) HandleGetFileServerHits(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(200)
	hits := handler.metrics.FileserverHits.Load()
	responseTemplate := `<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`
	rw.Write(fmt.Appendf(nil, responseTemplate, hits))
}

func (handler *MetricsHandler) MiddlewareFileServerHitsIncrement(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		handler.metrics.FileserverHits.Add(1)
		next.ServeHTTP(rw, req)
	})
}

func (handler *MetricsHandler) HandleResetMetricsReq(rw http.ResponseWriter, req *http.Request) {
	handler.metrics.FileserverHits.Store(0)
	rw.Header().Add("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(200)
	rw.Write([]byte("OK"))
}
